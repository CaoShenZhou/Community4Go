package routers

import (
	"github.com/CaoShenZhou/Blog4Go/internal/api"
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine) *gin.Engine {
	e.POST("/login", api.User.Login)
	e.POST("/reg", api.User.Reg)
	e.POST("/ValidateRegEmail", api.User.ValidateRegEmail)
	e.POST("/updatePwd", api.User.UpdatePwd)

	return e
}
