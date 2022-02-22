package user

import (
	"fmt"

	dao "github.com/CaoShenZhou/Blog4Go/internal/dao"
	model "github.com/CaoShenZhou/Blog4Go/internal/model"
	"github.com/CaoShenZhou/Blog4Go/pkg/util"
)

type UserService struct{}

// 查询邮箱是否注册
func (service *UserService) IsExistEmail(email string) bool {
	return dao.User.GetUserInfoByEmail(email) == (model.User{})
}

// 核对用户信息
func (service *UserService) CheckUserInfo(email, password string) (bool, model.User) {
	// 通过邮箱获取用户信息
	userInfo := dao.User.GetUserInfoByEmail(email)
	if userInfo != (model.User{}) {
		// 比对密码
		key := userInfo.ID[0:18] + "Mr.Cao"
		fmt.Println(key)
		aesPwd := util.AesEncrypt(password, key)
		fmt.Println(userInfo, aesPwd)
		fmt.Println(util.AesDecrypt(aesPwd, key))
		// 如果密码正确
		if aesPwd == userInfo.Password {
			return true, userInfo
		}
	}
	return false, userInfo
}
