package dto

// 验证注册用户
type ValidateRegUser struct {
	Email   string `json:"email" validate:"required,email"`
	Captcha string `json:"captcha" validate:"required,len=6"`
}

// binding:"required"
