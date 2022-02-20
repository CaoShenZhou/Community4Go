package dao

import (
	"github.com/CaoShenZhou/Blog4Go/global"
	model "github.com/CaoShenZhou/Blog4Go/internal/model"
)

// 通过邮箱获取用户信息
func GetUserInfoByEmail(email string) model.User {
	var isExistUser model.User
	global.Datasource.Where("email = ?", email).Limit(1).First(&isExistUser)
	return isExistUser
}

// 核对用户信息
func CheckUserInfo(email, password string) model.User {
	var userInfo model.User
	global.Datasource.Where("email = ? AND password = ?", email, password).First(&userInfo)
	return userInfo
}
