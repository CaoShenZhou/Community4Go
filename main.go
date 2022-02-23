package main

import (
	"fmt"
	"time"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/internal/middleware"
	router "github.com/CaoShenZhou/Blog4Go/internal/routers"
	"github.com/CaoShenZhou/Blog4Go/pkg/setting"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 数据库
func setupDatasource(setting *setting.DatasourceSetting) error {
	s := "%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local"
	db, err := gorm.Open(setting.DriverName,
		fmt.Sprintf(s,
			setting.Username,
			setting.Password,
			setting.Host,
			setting.Port,
			setting.Database,
			setting.Charset,
		))
	if err != nil {
		return err
	}
	global.Datasource = db
	fmt.Println(global.Datasource)
	// 这里可能会有数据的设置
	return nil
}

// 缓存
func setupRedis(setting *setting.RedisSetting) error {
	option := redis.DialPassword(setting.Password)
	conn, err := redis.Dial("tcp", setting.Host+":"+setting.Port, option)
	if err != nil {
		return err
	}
	global.Redis = conn
	return nil
}

func init() {
	// 读取配置文件
	setting, err := setting.ReadSetting()
	if err != nil {
		fmt.Println("读取配置文件错误,错误详情:", err)
		return
	}
	err = setting.ReadConfigStruct("server", &global.ServerSetting)
	if err != nil {
		fmt.Println("读取server配置错误,错误详情:", err)
		return
	}
	err = setting.ReadConfigStruct("datasource", &global.DatasourceSetting)
	if err != nil {
		fmt.Println("读取datasource配置错误,错误详情:", err)
		return
	}
	err = setting.ReadConfigStruct("jwt", &global.JwtSetting)
	if err != nil {
		fmt.Println("读取jwt配置错误,错误详情:", err)
		return
	}
	global.JwtSetting.Expire *= time.Second
	err = setting.ReadConfigStruct("redis", &global.RedisSetting)
	if err != nil {
		fmt.Println("读取redis配置错误,错误详情:", err)
		return
	}
	err = setting.ReadConfigStruct("email", &global.EmailSetting)
	if err != nil {
		fmt.Println("读取email配置错误,错误详情:", err)
		return
	}

	// 安装
	err = setupDatasource(global.DatasourceSetting)
	if err != nil {
		fmt.Println("连接datasource错误,错误详情:", err)
		return
	}
	err = setupRedis(global.RedisSetting)
	if err != nil {
		fmt.Println("连接redis错误,错误详情:", err)
		return
	}
}

func main() {
	r := gin.Default()
	// 公开组
	publicGroup := r.Group("")
	{
		router.Public.InitPublicRouter(publicGroup)
	}
	// 私有组
	privateGroup := r.Group("")
	privateGroup.Use(middleware.JwtAuth())
	{
		router.User.InitUserRouter(privateGroup)
	}
	panic(r.Run(":" + global.ServerSetting.Port))
}
