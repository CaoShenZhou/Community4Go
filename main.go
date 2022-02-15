package main

import (
	"fmt"

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
	err = setting.ReadConfigStruct("Server", &global.ServerSetting)
	if err != nil {
		fmt.Println("读取Server配置错误,错误详情:", err)
		return
	}
	err = setting.ReadConfigStruct("Datasource", &global.DatasourceSetting)
	if err != nil {
		fmt.Println("读取Datasource配置错误,错误详情:", err)
		return
	}

	err = setupDatasource()
	if err != nil {
		fmt.Println("连接Datasource错误,错误详情:", err)
		return
	}
}

func main() {
	r := gin.Default()
	r = routers.LoadUser(r)
	r = routers.LoadArticleTag(r)
	panic(r.Run(":" + global.ServerSetting.Port))
}
