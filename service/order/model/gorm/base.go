package gorm

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel struct {
	Id        int32      `gorm:"primarykey"`
	CreatedAt *time.Time `gorm:"column:add_time;" json:"create_time"` //这个是指针，这边要搞清楚为什么用指针
	UpdatedAt time.Time  `gorm:"column:update_time;"`

	//AddTime   time.Time `db:"add_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}
