package configs

import "time"

type JWT struct {
	Secret string        // 密钥
	Issuer string        // 签发人
	Expire time.Duration // 过期时间
}
