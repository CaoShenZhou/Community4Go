package initialize

import (
	"fmt"

	"github.com/CaoShenZhou/Community4Go/configs"
	"github.com/CaoShenZhou/Community4Go/global"
	"github.com/CaoShenZhou/Community4Go/model/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 数据库
func Datasource(config configs.Datasource) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		config.Config,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DriverName:                config.DriverName, // 驱动名
		DSN:                       dsn,               // 数据源名称:data source name
		DefaultStringSize:         256,               // string 类型字段的默认长度
		DisableDatetimePrecision:  true,              // 禁用datetime精度,MySQL5.6之前的数据库不支持
		DontSupportRenameIndex:    true,              // 重命名索引时采用删除并新建的方式,MySQL5.7之前的数据库和MariaDB不支持重命名索引
		DontSupportRenameColumn:   true,              // 用change重命名列,MySQL8之前的数据库和MariaDB不支持重命名列
		SkipInitializeWithVersion: false,             // 根据当前MySQL版本自动配置
	}), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Info),         // 打印SQL
		NamingStrategy: schema.NamingStrategy{SingularTable: true}}) // 命名策略
	if err != nil {
		return err
	}
	global.DB = db
	autoMigrateModels(db)
	return nil
}

// 自动迁移模型
func autoMigrateModels(db *gorm.DB) {
	if err := db.AutoMigrate(
		user.User{},
		user.UserLoginLog{},
	); err != nil {
		fmt.Println("数据模型迁移失败")
	}
}
