package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CaoShenZhou/Blog4Go/global"
	dto "github.com/CaoShenZhou/Blog4Go/internal/dto/user"
	"github.com/CaoShenZhou/Blog4Go/internal/entity"
	vo "github.com/CaoShenZhou/Blog4Go/internal/vo/user"
	"github.com/CaoShenZhou/Blog4Go/pkg/response"
	"github.com/CaoShenZhou/Blog4Go/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

func LoadUser(e *gin.Engine) *gin.Engine {
	e.POST("/login", Login)
	e.POST("/reg", Reg)
	e.POST("/checkEmail", CheckEmail)
	e.POST("/updatePwd", UpdatePwd)

	return e
}

func Login(ctx *gin.Context) {
	// 绑定模型
	var loginJson dto.LoginUser
	if err := ctx.ShouldBindJSON(&loginJson); err != nil {
		ctx.JSON(http.StatusBadRequest, response.BadRequest.WithMsg(err.Error()))
		return
	}
	// 校检模型数据
	validate := validator.New()
	err := validate.Struct(loginJson)
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
	global.Datasource.Debug().Where("email = ?", loginJson.Email).First(&userInfo)
	// 比对密码
	key := userInfo.ID[0:18] + "Mr.Cao"
	aesPwd := util.AesEncrypt(loginJson.Password, key)
	if aesPwd == userInfo.Password {
		// 拷贝模型字段
		var userInfoVo vo.LoginUser
		copier.Copy(&userInfoVo, &userInfo)
		// 生成token
		if token, err1 := util.GenerateToken(&userInfoVo); err1 == nil {
			vo := map[string]interface{}{
				"token":    token,
				"userInfo": userInfoVo,
			}
			ctx.JSON(http.StatusOK, response.Ok.WithData(vo))
		} else {
			fmt.Println(err1.Error())
			ctx.JSON(http.StatusInternalServerError, response.ServerError.WithMsg("token生成失败"))
			return
		}
	} else {
		ctx.JSON(http.StatusOK, response.InvalidParams.WithMsg("密码错误"))
	}
}
func Reg(c *gin.Context) {
	uuid := util.GetUUID()
	email := "caoshenzhou@gmail.com"
	pwd := "123456"
	key := uuid[0:18] + "Mr.Cao"
	AesPwd := util.AesEncrypt(pwd, key)
	user := entity.User{
		ID:        uuid,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: time.Now(),
		Email:     email,
		Nickname:  "张三",
		Password:  AesPwd,
	}
	// u1 := entity.User{
	// 	Nickname: "张三",
	// 	Password: "123456",
	// }
	// u2 := entity.User{
	// 	Nickname: "李四",
	// 	Password: "456789",
	// }
	// uList := []entity.User{
	// 	u1,
	// 	u2,
	// }
	// res := global.Datasource.NewRecord(user)
	// var userList []model.User
	// res := global.Datasource.Where("id = 123").Find(&userList)
	global.Datasource.Create(&user)
	c.JSON(http.StatusOK, response.Ok.WithData(user))
}
func CheckEmail(c *gin.Context) {
}
func UpdatePwd(c *gin.Context) {
}
