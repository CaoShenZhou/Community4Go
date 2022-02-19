package dto

// 注册用户
type RegUser struct {
	Nickname string `json:"nickname" validate:"required,min=2,max=18"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}
