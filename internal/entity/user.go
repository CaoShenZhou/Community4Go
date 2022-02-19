package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"deleted_at"`
	Email     string         `json:"email" gorm:"email,index"`
	Nickname  string         `json:"nickname" gorm:"nickname"`
	Password  string         `json:"password"  gorm:"password"`
}

func (u User) TableName() string {
	return "user"
}
