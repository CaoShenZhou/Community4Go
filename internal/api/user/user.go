package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CaoShenZhou/Blog4Go/global"
	model "github.com/CaoShenZhou/Blog4Go/internal/model"
	"github.com/CaoShenZhou/Blog4Go/internal/service"
	"github.com/CaoShenZhou/Blog4Go/pkg/response"
	"github.com/CaoShenZhou/Blog4Go/pkg/util"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type UserApi struct{}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

// 登录
func (api *UserApi) Login(ctx *gin.Context) {
	// 绑定模型
	var loginReq LoginReq
	fmt.Println(loginReq)
	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	if err := validator.New().Struct(loginReq); err != nil {
		/*for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}*/
		ctx.JSON(http.StatusPreconditionFailed, response.InvalidParams.WithMsg(err.Error()))
		return
	}
	// 查询用户是否存在
	if isExist, userInfo := service.User.CheckUserInfo(loginReq.Email, loginReq.Password); isExist {
		// 拷贝模型字段
		showUserInfo := gin.H{
			"id":       userInfo.ID,
			"nickname": userInfo.Nickname,
			"password": userInfo.Password,
		}
		copier.Copy(&showUserInfo, &userInfo)
		// 生成token
		if token, err := util.GenerateToken(showUserInfo); err == nil {
			var key = "user:token:" + userInfo.ID
			global.Redis.Do("set", key, token)
			global.Redis.Do("expire", key, 24*60*60)
			vo := gin.H{
				"token":    token,
				"userInfo": showUserInfo,
			}
			ctx.JSON(http.StatusOK, response.Ok.WithData(vo))
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, response.ServerError.WithMsg("token生成失败"))
			return
		}
	} else {
		ctx.JSON(http.StatusOK, response.InvalidParams.WithMsg("密码错误"))
		return
	}
}

type RegReq struct {
	Nickname string `json:"nickname" validate:"required,min=2,max=18"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

// 注册
func (api *UserApi) Reg(ctx *gin.Context) {
	// 绑定模型
	var req RegReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	if err := validator.New().Struct(req); err != nil {
		/*for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}*/
		ctx.JSON(http.StatusPreconditionFailed, response.InvalidParams.WithMsg(err.Error()))
		return
	}
	// 查询邮箱是否注册
	if service.User.IsExistEmail(req.Email) {
		to := []string{req.Email}
		captcha := util.RandomString(6, nil)
		if err := util.SendTextEmail(to, "注册博客", "您的验证码为："+captcha+"，有效期为五分钟"); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.EmailError)
		}
		key := "user:reg:userInfo:" + req.Email
		regUserJson, _ := json.Marshal(&req)
		global.Redis.Do("set", key, regUserJson)
		global.Redis.Do("expire", key, 5*60)
		key = "user:reg:captcha:" + req.Email
		global.Redis.Do("set", key, captcha)
		global.Redis.Do("expire", key, 5*60)
		ctx.JSON(http.StatusOK, response.Ok.WithMsg("验证码已发送，有效期为五分钟"))
		return
	} else {
		ctx.JSON(http.StatusOK, response.Unavailable.WithMsg("该邮箱已被注册"))
		return
	}
}

// 验证注册用户
type ValidateRegEmailReq struct {
	Email   string `json:"email" validate:"required,email"`
	Captcha string `json:"captcha" validate:"required,len=6"`
}

// 验证注册邮箱
func (api *UserApi) ValidateRegEmail(ctx *gin.Context) {
	// 绑定模型
	var req ValidateRegEmailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	if err := validator.New().Struct(req); err != nil {
		/*for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}*/
		ctx.JSON(http.StatusPreconditionFailed, response.InvalidParams.WithMsg(err.Error()))
		return
	}
	// 验证码key
	captchaKey := "user:reg:captcha:" + req.Email
	if captcha, err := redis.String(global.Redis.Do("get", captchaKey)); err != nil {
		if captcha == "" {
			ctx.JSON(http.StatusOK, response.ValidateFail.WithMsg("验证码失效"))
		} else {
			ctx.JSON(http.StatusInternalServerError, response.RedisError)
		}
	} else {
		// 如果验证码正确
		if captcha == req.Captcha {
			// 用户信息key
			userInfoKey := "user:reg:userInfo:" + req.Email
			if userInfoJson, err := redis.Bytes(global.Redis.Do("get", userInfoKey)); err != nil {
				ctx.JSON(http.StatusInternalServerError, response.RedisError)
				return
			} else {
				// 删除验证码缓存
				global.Redis.Do("del", captchaKey)
				// 删除用户信息缓存
				global.Redis.Do("del", userInfoKey)
				// 反序列化用户信息
				userInfo := &RegReq{}
				json.Unmarshal(userInfoJson, userInfo)
				// 填充用户信息
				uuid := util.GetUUID()
				pwd := userInfo.Password
				key := uuid[0:18] + "Mr.Cao"
				AesPwd := util.AesEncrypt(pwd, key)
				user := model.User{
					ID:       uuid,
					Email:    userInfo.Email,
					Nickname: userInfo.Nickname,
					Password: AesPwd,
				}
				// 写入数据
				global.Datasource.Create(&user)
				ctx.JSON(http.StatusOK, response.Ok.WithMsg("注册成功"))
				return
			}
		} else {
			ctx.JSON(http.StatusOK, response.ValidateFail.WithMsg("验证码错误"))
			return
		}
	}
}

// 更新密码
func (api *UserApi) UpdatePwd(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.Ok)
	return
}
