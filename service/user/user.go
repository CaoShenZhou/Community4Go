package user

import (
	"fmt"

	"github.com/CaoShenZhou/Blog4Go/dao"
	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/model/user"
	"github.com/CaoShenZhou/Blog4Go/util"
)

type UserService struct{}

// 查询用户名是否已被注册
func (service *UserService) IsUsernameExists(usernameType, username string) (bool, error) {
	db := global.DB.Model(&user.User{})
	if usernameType == user.UsernameTypeEmail {
		db = db.Where("email = ?", username)
	}
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

// 注册用户
func (service *UserService) Register(loginMethod string, user user.User) error {
	if err := global.DB.Create(&user).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// 按用户名获取用户信息
func (service *UserService) Login(usernameType, username, pwd string) (*user.User, error) {
	if userInfo, err := dao.User.GetUserInfoByUsername(usernameType, username); err != nil {
		return nil, err
	} else {
		if userInfo == nil {
			return nil, nil
		} else {
			// 比对密码
			userID := fmt.Sprintf("%d", userInfo.ID)
			key := userID[len(userID)-10:] + "Mr.Cao"
			aesPwd := util.AESEncrypt(pwd, key)
			// 如果密码正确
			if aesPwd == userInfo.Password {
				return userInfo, nil
			} else {
				return nil, nil
			}
		}
	}
}
