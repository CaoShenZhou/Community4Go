package user

import (
	"github.com/CaoShenZhou/Blog4Go/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (router *UserRouter) PrivateUserRouter(rg *gin.RouterGroup) {
	// baseRouter := rg.Group("user")
	// baseApi := api.User
	{
	}
}

func (router *UserRouter) PublicUserRouter(rg *gin.RouterGroup) {
	baseRouter := rg.Group("user")
	baseApi := api.User
	{
		baseRouter.POST("/login", baseApi.Login)
		baseRouter.POST("/getRegisterCaptcha", baseApi.GetRegisterCaptcha)
		baseRouter.POST("/register", baseApi.Register)
	}
}
