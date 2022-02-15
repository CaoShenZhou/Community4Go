package entity

import "time"

type User struct {
	ID        string    `gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Email     string    `json:"email" gorm:"email"`
	Nickname  string    `json:"nickname" gorm:"nickname"`
	Password  string    `json:"password"  gorm:"password"`
}

func (u User) TableName() string {
	return "user"
}
