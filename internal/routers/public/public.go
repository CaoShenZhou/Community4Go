package public

import (
	"github.com/CaoShenZhou/Blog4Go/internal/api"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct{}

func (router *PublicRouter) InitPublicRouter(rg *gin.RouterGroup) {
	baseRouter := rg.Group("")
	{
		baseRouter.POST("/login", api.User.Login)
		baseRouter.POST("/reg", api.User.Reg)
		baseRouter.POST("/ValidateRegEmail", api.User.ValidateRegEmail)
	}
}
