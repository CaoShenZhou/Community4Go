package model

import (
	"fmt"

	"github.com/CaoShenZhou/Blog4Go/pkg/setting"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDatasource(setting *setting.DatasourceSetting) (*gorm.DB, error) {
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
		return db, nil
	}
	// 这里可能会有数据的设置
	return db, nil
}
