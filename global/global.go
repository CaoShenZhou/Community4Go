package global

import (
	"github.com/CaoShenZhou/Community4Go/configs"
	"github.com/garyburd/redigo/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Redis  redis.Conn
	Email  configs.Email
	JWT    configs.JWT
	Server configs.Server
	Zap    zap.Logger
)
