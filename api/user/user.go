package user

import (
	"fmt"
	"net/http"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/model/response"
	"github.com/CaoShenZhou/Blog4Go/model/user"
	"github.com/CaoShenZhou/Blog4Go/service"
	"github.com/CaoShenZhou/Blog4Go/util"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// 登录
func (api *UserApi) Login(c *gin.Context) {
	// 请求参数
	type reqParm struct {
		UsernameType string `json:"username_type" validate:"required,oneof=Email MSISDN"` // 用户名类型
		Username     string `json:"Username" validate:"required"`                         // 用户名
		Password     string `json:"password" validate:"required,min=6,max=32"`            // 密码
	}
	// 绑定参数
	var req reqParm
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.PARM_BIND_FAIL)
		return
	}
	// 校检参数
	if err := util.VerifyParm(req); err != nil {
		response.FailWithMsg(c, response.PARM_VERIFY_FAIL)
		return
	}
	// 查询用户是否存在
	if userInfo, err := service.User.Login(user.UsernameTypeEmail, req.Username, req.Password); err != nil {
		response.Error(c, http.StatusInternalServerError)
		return
	} else {
		if userInfo == nil {
			response.FailWithMsg(c, "账号密码不匹配")
			return
		}
		// 拷贝模型字段
		showUserInfo := gin.H{
			"id":       userInfo.ID,
			"nickname": userInfo.Nickname,
		}
		// 生成token
		if token, err := util.GenerateToken(showUserInfo); err == nil {
			var key = "user:token:" + fmt.Sprintf("%d", userInfo.ID)
			global.Redis.Do("set", key, token)
			global.Redis.Do("expire", key, 24*60*60) // 过期时间为一天
			vo := gin.H{
				"token":    token,
				"userInfo": showUserInfo,
			}
			response.OkWithMsgAndData(c, "登录成功", vo)
			return
		} else {
			response.Error(c, http.StatusInternalServerError)
			return
		}
	}
}

// 获取注册验证码
func (api *UserApi) GetRegisterCaptcha(c *gin.Context) {
	type reqParm struct {
		UsernameType string `json:"username_type" validate:"required,oneof=Email MSISDN"` // 用户名类型
		Username     string `json:"Username" validate:"required"`                         // 用户名
	}
	// 绑定参数
	var req reqParm
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.PARM_BIND_FAIL)
		return
	}
	// 校检参数
	if err := util.VerifyParm(req); err != nil {
		response.FailWithMsg(c, response.PARM_VERIFY_FAIL)
		return
	}
	// 查询用户名是否已被注册
	if isExists, err := service.User.IsUsernameExists(req.UsernameType, req.Username); err != nil {
		response.Error(c, http.StatusInternalServerError)
		return
	} else {
		vo := gin.H{}
		vo["isExists"] = isExists
		if isExists {
			response.OkWithData(c, vo)
			return
		} else {
			captcha := util.RandomString(6, nil)
			vo["captcha"] = captcha
			to := []string{req.Username}
			if err := util.SendTextEmail(to, "注册博客", "您的验证码为："+captcha+"，有效期为五分钟"); err != nil {
				response.Error(c, http.StatusInternalServerError)
			} else {
				response.OkWithData(c, "验证码已发送")
			}
			return
		}
	}
}

/*
// 注册
func (api *UserApi) Register(c *gin.Context) {

	sf, err := util.GetSnowflake()
	if err == nil {
		sfStr := sf.String()
		pwd := "123456"
		key := sfStr[len(sfStr)-10:] + "Mr.Cao"
		fmt.Printf("密钥%s\n", key)
		password := util.AESEncrypt(pwd, key)
		a := uint(sf.Int64())
		user := model.User{
			BaseModel:   model.BaseModel{ID: a},
			Email:       "123@gmail.com",
			PhoneNumber: "16600001111",
			Password:    password,
		}
		err := global.DB.Create(&user)
		fmt.Println(err)
	}
	type reqParm struct {
		Nickname string `json:"nickname" validate:"required,min=2,max=18"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=32"`
	}
	// 绑定参数
	var req reqParm
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.PARM_BIND_FAIL)
		return
	}
	// 校检参数
	if err := util.VerifyParm(req); err != nil {
		response.FailWithMsg(c, response.PARM_VERIFY_FAIL)
		return
	}
	// 查询邮箱是否注册
	if service.User.IsExistEmail("email", req.Email) {
		to := []string{req.Email}
		captcha := util.RandomString(6, nil)
		if err := util.SendTextEmail(to, "注册博客", "您的验证码为："+captcha+"，有效期为五分钟"); err != nil {
			response.Error(c, http.StatusInternalServerError)
		}
		key := "user:reg:userInfo:" + req.Email
		regUserJson, _ := json.Marshal(&req)
		global.Redis.Do("set", key, regUserJson)
		global.Redis.Do("expire", key, 5*60)
		key = "user:reg:captcha:" + req.Email
		global.Redis.Do("set", key, captcha)
		global.Redis.Do("expire", key, 5*60)
		response.FailWithMsg(c, "验证码已发送，有效期为五分钟")
		return
	} else {
		response.FailWithMsg(c, "该邮箱已被注册")
		return
	}
}

/*
// 验证注册邮箱
func (api *UserApi) VerifyRegEmail(c *gin.Context) {
	// 验证注册用户
	type reqParm struct {
		Email   string `json:"email" validate:"required,email"`
		Captcha string `json:"captcha" validate:"required,len=6"`
	}
	// 绑定参数
	var req reqParm
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMsg(c, response.PARM_BIND_FAIL)
		return
	}
	// 校检参数
	if err := util.VerifyParm(req); err != nil {
		response.FailWithMsg(c, response.PARM_VERIFY_FAIL)
		return
	}
	// 验证码key
	captchaKey := "user:reg:captcha:" + req.Email
	if captcha, err := redis.String(global.Redis.Do("get", captchaKey)); err != nil {
		if captcha == "" {
			response.FailWithMsg(c, "验证码失效")
			return
		} else {
			response.Error(c, http.StatusInternalServerError)
			return
		}
	} else {
		// 如果验证码正确
		if captcha == req.Captcha {
			// 用户信息key
			userInfoKey := "user:reg:userInfo:" + req.Email
			if userInfoJson, err := redis.Bytes(global.Redis.Do("get", userInfoKey)); err != nil {
				response.Error(c, http.StatusInternalServerError)
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
				uuid := "123"
				pwd := userInfo.Password
				key := uuid[0:18] + "Mr.Cao"
				AesPwd := util.AesEncrypt(pwd, key)
				user := model.User{
					Email:    userInfo.Email,
					Nickname: userInfo.Nickname,
					Password: AesPwd,
				}
				// 写入数据
				global.DB.Create(&user)
				response.OkWithMsg(c, "注册成功")
				return
			}
		} else {
			response.FailWithMsg(c, "验证码错误")
			return
		}
	}
}
*/
