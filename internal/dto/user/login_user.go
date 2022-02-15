package dto

// 登录用户
type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=18"`
}

// binding:"required"
