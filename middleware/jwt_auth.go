package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CaoShenZhou/Community4Go/model/user"
	"github.com/CaoShenZhou/Community4Go/util"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		fmt.Println(path)
		// 跳到下一个中间件
		// c.Next()
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, "未认证")
			// 中间件结束后会直接返回
			c.Abort()
			return
		} else {
			if claims, err := util.ParseToken(token); err != nil {
				c.JSON(http.StatusUnauthorized, "未认证")
				c.Abort()
				return
			} else {
				tokenInfo := user.UserTokenInfo{}
				json.Unmarshal([]byte(claims.Info), &tokenInfo)
				fmt.Println(tokenInfo)
				return
			}
		}
	}
}
