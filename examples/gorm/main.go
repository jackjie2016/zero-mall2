package main

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"zero-mal/examples/gorm/goods/model"
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

type BaseModel1 struct {
	ID        int32      `gorm:"primarykey"`
	CreatedAt *time.Time `gorm:"column:add_time"`
	UpdatedAt time.Time  `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt
	IsDeleted bool
}
type (
	BaseModel struct {
		Id        int32      `gorm:"primarykey"` //为什么使用int32， bigint
		CreatedAt *time.Time `gorm:"column:add_time"`
		UpdatedAt time.Time  `gorm:"column:update_time"`
		DeletedAt gorm.DeletedAt
		IsDeleted bool
	}

	Category1 struct {
		BaseModel
		Name             string      `gorm:"type:varchar(20);not null" json:"name"`
		ParentCategoryID int32       `json:"parent"`
		ParentCategory   *Category   `json:"-"`
		SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
		Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
		IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
	}

	Category struct {
		BaseModel
		Name             string      `gorm:"type:varchar(20);not null" json:"name"`
		ParentCategoryID int32       `json:"parent"`
		ParentCategory   *Category   `json:"-"`
		SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:Id" json:"sub_category"`
		Level            int32       `gorm:"type:int;not null;default:1" json:"level"`
		IsTab            bool        `gorm:"default:false;not null" json:"is_tab"`
	}
)

func main() {

	dsn := "root:root123@tcp(127.0.0.1:3306)/go-zero-mall?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "mxshop_",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	getCate(db)
	getCate1(db)
}
func getCate(db *gorm.DB) {
	var categorys []model.Category
	db.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	fmt.Printf("%+v", categorys)

}

func getCate1(db *gorm.DB) {
	var categorys []Category
	db.Where(&Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	fmt.Printf("%+v", categorys)

}
