package routers

import (
	"encoding/json"
	"net/http"

	"github.com/CaoShenZhou/Blog4Go/global"
	dto "github.com/CaoShenZhou/Blog4Go/internal/dto/user"
	"github.com/CaoShenZhou/Blog4Go/internal/entity"
	vo "github.com/CaoShenZhou/Blog4Go/internal/vo/user"
	"github.com/CaoShenZhou/Blog4Go/pkg/response"
	"github.com/CaoShenZhou/Blog4Go/pkg/util"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

func LoadUser(e *gin.Engine) *gin.Engine {
	e.POST("/login", Login)
	e.POST("/reg", Reg)
	e.POST("/ValidateRegEmail", ValidateRegEmail)
	e.POST("/updatePwd", UpdatePwd)

	return e
}

// 登录
func Login(ctx *gin.Context) {
	// 绑定模型
	var loginUser dto.LoginUser
	if err := ctx.ShouldBindJSON(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	validate := validator.New()
	err := validate.Struct(loginUser)
	if err != nil {
		/*for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}*/
		ctx.JSON(http.StatusPreconditionFailed, response.InvalidParams.WithMsg(err.Error()))
		return
	}
	// 查询用户是否存在
	var userInfo entity.User
	global.Datasource.Where("email = ?", loginUser.Email).First(&userInfo)
	// 比对密码
	key := userInfo.ID[0:18] + "Mr.Cao"
	aesPwd := util.AesEncrypt(loginUser.Password, key)
	if aesPwd == userInfo.Password {
		// 拷贝模型字段
		var userInfoVo vo.LoginUser
		copier.Copy(&userInfoVo, &userInfo)
		// 生成token
		if token, err := util.GenerateToken(&userInfoVo); err == nil {
			var key = "token:user:" + userInfo.ID
			global.Redis.Do("set", key, token)
			global.Redis.Do("expire", key, 24*60*60)
			vo := map[string]interface{}{
				"token":    token,
				"userInfo": userInfoVo,
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

// 注册
func Reg(ctx *gin.Context) {
	// 绑定模型
	var regUser dto.RegUser
	if err := ctx.ShouldBindJSON(&regUser); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	validate := validator.New()
	err := validate.Struct(regUser)
	if err != nil {
		/*for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}*/
		ctx.JSON(http.StatusPreconditionFailed, response.InvalidParams.WithMsg(err.Error()))
		return
	}
	// 查询邮箱是否注册
	var isExistUser entity.User
	global.Datasource.Where("email = ?", regUser.Email).Limit(1).First(&isExistUser)
	if isExistUser == (entity.User{}) {
		to := []string{regUser.Email}
		captcha := util.RandomString(6, nil)
		if err := util.SendTextEmail(to, "注册博客", "您的验证码为："+captcha+"，有效期为五分钟"); err != nil {
			ctx.JSON(http.StatusInternalServerError, response.EmailError)
		}
		key := "user:reg:userInfo:" + regUser.Email
		regUserJson, _ := json.Marshal(&regUser)
		global.Redis.Do("set", key, regUserJson)
		global.Redis.Do("expire", key, 5*60)
		key = "user:reg:captcha:" + regUser.Email
		global.Redis.Do("set", key, captcha)
		global.Redis.Do("expire", key, 5*60)
		ctx.JSON(http.StatusOK, response.Ok.WithMsg("验证码已发送，有效期为五分钟"))
		return
	} else {
		ctx.JSON(http.StatusOK, response.Unavailable.WithMsg("该邮箱已被注册"))
		return
	}
}

// 验证注册邮箱
func ValidateRegEmail(ctx *gin.Context) {
	// 绑定模型
	var vru dto.ValidateRegUser
	if err := ctx.ShouldBindJSON(&vru); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	validate := validator.New()
	err := validate.Struct(vru)
	if err != nil {
		/*for _, e := range err.(validator.ValidationErrors) {
			fmt.Println("错误字段：", e.Field())
			fmt.Println("错误值：", e.Value())
			fmt.Println("错误tag：", e.Tag())
		}*/
		ctx.JSON(http.StatusPreconditionFailed, response.InvalidParams.WithMsg(err.Error()))
		return
	}
	// 验证码key
	captchaKey := "user:reg:captcha:" + vru.Email
	if captcha, err := redis.String(global.Redis.Do("get", captchaKey)); err != nil {
		if captcha == "" {
			ctx.JSON(http.StatusInternalServerError, response.ValidateFail.WithMsg("验证码失效"))
		} else {
			ctx.JSON(http.StatusInternalServerError, response.RedisError)
		}
	} else {
		// 如果验证码正确
		if captcha == vru.Captcha {
			// 用户信息key
			userInfoKey := "user:reg:userInfo:" + vru.Email
			if userInfoJson, err := redis.Bytes(global.Redis.Do("get", userInfoKey)); err != nil {
				ctx.JSON(http.StatusInternalServerError, response.RedisError)
				return
			} else {
				// 删除验证码缓存
				global.Redis.Do("del", captchaKey)
				// 删除用户信息缓存
				global.Redis.Do("del", userInfoKey)
				// 反序列化用户信息
				userInfo := &dto.RegUser{}
				json.Unmarshal(userInfoJson, userInfo)
				// 填充用户信息
				uuid := util.GetUUID()
				nickname := userInfo.Nickname
				email := userInfo.Email
				pwd := userInfo.Password
				key := uuid[0:18] + "Mr.Cao"
				AesPwd := util.AesEncrypt(pwd, key)
				user := entity.User{
					ID:       uuid,
					Email:    email,
					Nickname: nickname,
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
func UpdatePwd(c *gin.Context) {
}
