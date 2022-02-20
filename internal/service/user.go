package service

import (
	"github.com/CaoShenZhou/Blog4Go/internal/dao"
	model "github.com/CaoShenZhou/Blog4Go/internal/model"
	"github.com/CaoShenZhou/Blog4Go/pkg/util"
)

// 查询邮箱是否注册
func IsExistEmail(email string) bool {
	return dao.GetUserInfoByEmail(email) == (model.User{})
}

// 核对用户信息
func CheckUserInfo(email, password string) (bool, model.User) {
	// 通过邮箱获取用户信息
	userInfo := dao.GetUserInfoByEmail(email)
	// 比对密码
	key := userInfo.ID[0:18] + "Mr.Cao"
	aesPwd := util.AesEncrypt(password, key)
	// 如果密码正确
	if aesPwd == userInfo.Password {
		return true, userInfo
	}
	return false, userInfo
}
