package user

import "github.com/CaoShenZhou/Blog4Go/model"

type User struct {
	model.BaseModel
	Email    string `json:"email" gorm:"index"`  // 电子邮箱
	MSISDN   string `json:"msisdn" gorm:"index"` // 全球唯一手机号
	Nickname string `json:"nickname"`            // 昵称
	Password string `json:"password"`            // 密码
}

const (
	UsernameTypeEmail  = "Email"
	UsernameTypeMSISDN = "MSISDN"
)

// 用户令牌信息
type UserTokenInfo struct {
	UserID   uint   `json:"user_id"`  // 用户ID
	Nickname string `json:"nickname"` // 昵称
}
