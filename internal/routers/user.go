package routers

import (
	"net/http"

	"github.com/CaoShenZhou/Blog4Go/pkg/response"
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine) *gin.Engine {
	e.POST("/login", Login)
	e.POST("/reg", Reg)
	e.POST("/checkEmail", CheckEmail)
	e.POST("/updatePwd", UpdatePwd)

	return e
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, response.Ok)
}
func Reg(c *gin.Context) {
}
func CheckEmail(c *gin.Context) {
}
func UpdatePwd(c *gin.Context) {
}
