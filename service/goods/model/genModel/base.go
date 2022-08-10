package genModel

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int64     `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"column:create_time"`
	UpdatedAt time.Time `gorm:"column:update_time"`

	AddTime   time.Time `db:"add_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}
