package user

import (
	"github.com/CaoShenZhou/Blog4Go/global"
	model "github.com/CaoShenZhou/Blog4Go/internal/model"
)

type UserDao struct{}

// 通过邮箱获取用户信息
func (dao *UserDao) GetUserInfoByEmail(email string) model.User {
	var userInfo model.User
	global.Datasource.Where("email = ?", email).Limit(1).First(&userInfo)
	return userInfo
}

// 核对用户信息
func (dao *UserDao) CheckUserInfo(email, password string) model.User {
	var userInfo model.User
	global.Datasource.Where("email = ? AND password = ?", email, password).First(&userInfo)
	return userInfo
}
