package main

import (
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
}
func getCate(db *gorm.DB) {
	var categorys []model.Category
	db.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	fmt.Printf("%+v", categorys)

}
