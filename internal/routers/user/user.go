package user

import (
	"github.com/CaoShenZhou/Blog4Go/internal/api"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (router *UserRouter) InitUserRouter(rg *gin.RouterGroup) {
	baseRouter := rg.Group("user")
	baseApi := api.User
	{
		baseRouter.POST("/updatePwd", baseApi.UpdatePwd)
	}
}
