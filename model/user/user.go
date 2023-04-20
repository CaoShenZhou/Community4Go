package user

import "github.com/CaoShenZhou/Community4Go/model"

type User struct {
	model.BaseModel
	Username       string `json:"username" gorm:"uniqueIndex"` // 用户名
	Email          string `json:"email" gorm:"uniqueIndex"`    // 电子邮箱
	IDD            string `json:"idd"`                         // 国际区号
	MobileNumber   string `json:"mobile_number"`               // 手机号码
	MSISDN         string `json:"msisdn" gorm:"uniqueIndex"`   // 全球唯一手机号
	ProfilePicture string `json:"profile_picture"`             // 头像
	Nickname       string `json:"nickname"`                    // 昵称
	Password       string `json:"password"`                    // 密码
	Lang           string `json:"lang"`                        // 语言
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
