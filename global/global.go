package global

import (
	"github.com/CaoShenZhou/Blog4Go/configs"
	"github.com/garyburd/redigo/redis"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Redis  redis.Conn
	Email  configs.Email
	JWT    configs.JWT
	Server configs.Server
)
