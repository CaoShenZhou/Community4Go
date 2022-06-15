package main

import (
	"fmt"
	"time"

	"github.com/CaoShenZhou/Blog4Go/configs"
	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/initialize"
	"github.com/CaoShenZhou/Blog4Go/middleware"
	router "github.com/CaoShenZhou/Blog4Go/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	// 读取配置文件
	setting, err := initialize.ReadConfig()
	if err != nil {
		fmt.Println("读取配置文件错误,错误详情:", err)
		return
	}
	// 读取服务配置
	if err := setting.ReadConfigStruct("server", &global.Server); err != nil {
		fmt.Println("读取server配置错误,错误详情:", err)
		return
	}
	// 配置JWT
	if err := setting.ReadConfigStruct("jwt", &global.JWT); err != nil {
		fmt.Println("读取jwt配置错误,错误详情:", err)
		return
	}
	global.JWT.Expire *= time.Second
	// 配置邮箱
	if err := setting.ReadConfigStruct("email", &global.Email); err != nil {
		fmt.Println("读取email配置错误,错误详情:", err.Error())
		return
	}
	// 初始化数据库
	var configDB configs.Datasource
	if err := setting.ReadConfigStruct("datasource", &configDB); err != nil {
		fmt.Println("读取datasource配置错误,错误详情:", err)
		return
	} else {
		if err := initialize.Datasource(configDB); err != nil {
			fmt.Println("连接datasource错误,错误详情:", err.Error())
			return
		}
	}
	// 初始化缓存
	var configRedis configs.Redis
	if err := setting.ReadConfigStruct("redis", &configRedis); err != nil {
		fmt.Println("连读取redis配置错误,错误详情:", err.Error())
		return
	} else {
		if err := initialize.Redis(configRedis); err != nil {
			fmt.Println("连接redis错误,错误详情:", err.Error())
			return
		}
	}

}

func main() {
	r := gin.Default()
	// 公开组
	publicGroup := r.Group("")
	{
		router.User.PublicUserRouter(publicGroup)
	}
	// 私有组
	privateGroup := r.Group("")
	privateGroup.Use(middleware.JwtAuth())
	{
		router.User.PrivateUserRouter(privateGroup)
	}
	panic(r.Run(fmt.Sprintf(":%v", global.Server.Port)))
}
