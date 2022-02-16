package main

import (
	"fmt"
	"time"

	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/internal/model"
	"github.com/CaoShenZhou/Blog4Go/internal/routers"
	"github.com/CaoShenZhou/Blog4Go/pkg/setting"
	"github.com/gin-gonic/gin"
)

// 连接数据库
func setupDatasource() error {
	var err error
	global.Datasource, err = model.NewDatasource(global.DatasourceSetting)
	if err != nil {
		return err
	}
	return err
}

func init() {
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
		fmt.Println("读取Jwt配置错误,错误详情:", err)
		return
	}
	global.JwtSetting.Expire *= time.Second

	err = setupDatasource()
	if err != nil {
		fmt.Println("连接datasource错误,错误详情:", err)
		return
	}
}

func main() {
	r := gin.Default()
	r = routers.LoadUser(r)
	r = routers.LoadArticleTag(r)
	panic(r.Run(":" + global.ServerSetting.Port))
}
