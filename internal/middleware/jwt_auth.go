package middleware

import (
	"fmt"
	"net/http"

	"github.com/CaoShenZhou/Blog4Go/pkg/response"
	"github.com/CaoShenZhou/Blog4Go/pkg/util"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		fmt.Println(path)
		// 跳到下一个中间件
		// ctx.Next()
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, response.Unauthenticated)
			// 中间件结束后会直接返回
			ctx.Abort()
			return
		} else {
			if claims, err := util.ParseToken(token); err != nil {
				ctx.JSON(http.StatusUnauthorized, response.Unauthenticated.WithMsg(err.Error()))
				ctx.Abort()
				return
			} else {
				fmt.Println(claims)
				return
			}
		}
	}
}
