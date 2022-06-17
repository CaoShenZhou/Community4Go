package user

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/model"
	"github.com/CaoShenZhou/Blog4Go/model/response"
	"github.com/CaoShenZhou/Blog4Go/model/user"
	"github.com/CaoShenZhou/Blog4Go/service"
	"github.com/CaoShenZhou/Blog4Go/util"
	"github.com/garyburd/redigo/redis"
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
		// 用户令牌信息
		userTokenInfo := user.UserTokenInfo{
			UserID:   userInfo.ID,
			Nickname: userInfo.Nickname,
		}
		jsonByteArr, _ := json.Marshal(userTokenInfo)
		// 生成token
		if token, err := util.GenerateToken(string(jsonByteArr)); err == nil {
			var key = "user:token:" + fmt.Sprintf("%d", userInfo.ID)
			global.Redis.Do("set", key, token)
			global.Redis.Do("expire", key, 24*60*60) // 过期时间为一天
			vo := gin.H{
				"token":    token,
				"userInfo": userTokenInfo,
			}
			ip, _ := c.RemoteIP()
			ull := user.UserLoginLog{
				UserID: userInfo.ID,
				IP:     ip.String(),
			}
			service.User.AddLoginLog(ull)
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
		}
		// 注册验证码的key
		key := fmt.Sprintf("user:reg:%s:%s", req.UsernameType, req.Username)
		// 如果存在验证码就不再次创建
		if value, err := redis.String(global.Redis.Do("get", key)); value == "" && err != redis.ErrNil {
			fmt.Println(err)
			fmt.Println(value)
			fmt.Println("[][][]")
			response.Error(c, http.StatusInternalServerError)
			return
		} else {
			if value != "" {
				response.OkWithMsg(c, "验证码过期后才可再次发送")
				return
			}
		}
		// 验证码
		rand.Seed(time.Now().UnixNano())
		captcha := util.RandomString(6, "")
		to := []string{req.Username}
		if err := util.SendTextEmail(to, "注册博客", "您的验证码为："+captcha+"，有效期为五分钟"); err != nil {
			response.Error(c, http.StatusInternalServerError)
		} else {
			global.Redis.Do("set", key, captcha)
			global.Redis.Do("expire", key, 5*60) // 5分钟
			response.OkWithMsg(c, "验证码已发送")
		}
		return
	}
}

// 注册
func (api *UserApi) Register(c *gin.Context) {
	// 请求参数
	type reqParm struct {
		Nickname     string `json:"nickname" validate:"required,min=2,max=18"`            // 昵称
		Captcha      string `json:"captcha" validate:"required,len=6"`                    // 验证码
		UsernameType string `json:"username_type" validate:"required,oneof=Email MSISDN"` // 用户名类型
		Username     string `json:"username" validate:"required"`                         // 用户名
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
	// 注册验证码的key
	key := fmt.Sprintf("user:reg:%s:%s", req.UsernameType, req.Username)
	if captcha, err := redis.String(global.Redis.Do("get", key)); captcha == "" && err != redis.ErrNil {
		fmt.Println("123")
		response.Error(c, http.StatusInternalServerError)
		return
	} else {
		if req.Captcha == captcha {
			// 开始注册流程
			if sf, err := util.GetSnowflake(); err != nil {
				response.Error(c, http.StatusInternalServerError)
				return
			} else {
				sfStr := sf.String()
				pwdKey := sfStr[len(sfStr)-10:] + "Mr.Cao"
				password := util.AESEncrypt(req.Password, pwdKey)
				userInfo := user.User{
					BaseModel: model.BaseModel{ID: uint(sf.Int64())},
					Nickname:  req.Nickname,
					Password:  password,
				}
				if req.UsernameType == user.UsernameTypeEmail {
					userInfo.Email = req.Username
				} else if req.UsernameType == user.UsernameTypeMSISDN {
					userInfo.MSISDN = req.Username
				}
				if err := global.DB.Create(&userInfo).Error; err != nil {
					response.Error(c, http.StatusInternalServerError)
					return
				}
				global.Redis.Do("del", key)
			}
			response.OkWithMsg(c, "注册成功")
		} else {
			response.FailWithMsg(c, "验证码有误")
		}
		return
	}
}
