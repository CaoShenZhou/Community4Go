package model

import "gorm.io/plugin/soft_delete"

type BaseModel struct {
	ID        uint                  `json:"id" gorm:"primarykey"`                          // 主键ID
	CreatedAt int64                 `json:"created_at" gorm:"index:,autoCreateTime:milli"` // 创建时间
	UpdatedAt int64                 `json:"updated_at" gorm:"index:,autoUpdateTime:milli"` // 更新时间
	DeletedAt soft_delete.DeletedAt `json:"-" gorm:"index:,softDelete:milli"`              // 删除时间
}
