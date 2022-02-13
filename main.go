package main

import (
	"github.com/CaoShenZhou/Blog4Go/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// id := util.GetUUID()
	// fmt.Println(id)
	// pwd := "12345"
	// // 获取前24位字符
	// key := id[0:24]
	// fmt.Println("原文：", pwd)
	// encryptCode := util.AesEncrypt(pwd, key)
	// fmt.Println("解密结果：", encryptCode)
	// decryptCode := util.AesDecrypt(encryptCode, key)
	// fmt.Println("解密结果：", decryptCode)

	r := gin.Default()
	r = routers.LoadUser(r)
	r = routers.LoadArticleTag(r)
	panic(r.Run(":8088"))
}
