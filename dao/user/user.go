package user

import (
	"github.com/CaoShenZhou/Blog4Go/global"
	"github.com/CaoShenZhou/Blog4Go/model/user"
	"gorm.io/gorm"
)

type UserDao struct{}

// 按照用户名获取用户信息
func (dao *UserDao) GetUserInfoByUsername(usernameType, username string) (*user.User, error) {
	db := global.DB.Model(&user.User{})
	if usernameType == user.UsernameTypeEmail {
		db = db.Where("email = ?", username)
	}
	var userInfo user.User
	if err := db.First(&userInfo).Error; err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		return &userInfo, nil
	}
}
