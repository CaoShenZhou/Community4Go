package user

import "github.com/CaoShenZhou/Blog4Go/model"

type UserLoginLog struct {
	model.BaseModel
	UserID uint   `json:"user_id" gorm:"index:"` // 用户ID
	IP     string `json:"ip"`                    // 登录IP
}
