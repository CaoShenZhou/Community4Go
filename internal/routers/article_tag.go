package routers

import "github.com/gin-gonic/gin"

func LoadArticleTag(e *gin.Engine) *gin.Engine {
	e.POST("/add", Add)
	e.DELETE("/del", Del)
	e.POST("/edit", Edit)

	return e
}

func Add(c *gin.Context) {
}
func Del(c *gin.Context) {
}
func Edit(c *gin.Context) {
}
