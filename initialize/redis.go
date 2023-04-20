package initialize

import (
	"github.com/CaoShenZhou/Community4Go/configs"
	"github.com/CaoShenZhou/Community4Go/global"
	"github.com/garyburd/redigo/redis"
)

// 缓存
func Redis(config configs.Redis) error {
	option := redis.DialPassword(config.Password)
	conn, err := redis.Dial("tcp", config.Host+":"+config.Port, option)
	if err != nil {
		return err
	}
	global.Redis = conn
	return nil
}
