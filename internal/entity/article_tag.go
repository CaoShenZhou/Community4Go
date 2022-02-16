package entity

import "time"

type ArticleTag struct {
	ID        string    `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"deleted_at"`
	Name      string    `json:"name" gorm:"name"`
}

func (at ArticleTag) TableName() string {
	return "article_tag"
}
