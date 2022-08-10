package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type (
	BaseModel struct {
		Id int32 `gorm:"primarykey"`

		CreatedAt time.Time `gorm:"column:add_time"`
		UpdatedAt time.Time `gorm:"column:update_time"`
		DeletedAt gorm.DeletedAt
		IsDeleted bool
	}

	//md5 信息摘要算法
	User struct {
		BaseModel
		Mobile   string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
		Password string     `gorm:"type:varchar(100);not null"`
		NickName string     `gorm:"type:varchar(20);not null"`
		Birthday *time.Time `gorm:"type:datetime"`
		Gender   string     `gorm:"column:gender;default:male;type:varchar(6) comment 'female 表示女 male表示男'"`
		Role     int        `gorm:"column:role;default:1;type:int(1) comment '1 表示普通用户 2表示管理员'"`
	}
	Product struct {
		gorm.Model
		Code  string
		Price uint
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

	// 迁移 schema
	//db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1)                 // 根据整型主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//
	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)
	var user User
	//db.Where("mobile", "15958615799").First(&user, 1) // 根据整型主键查找
	db.Where("mobile", "15958615799").First(&user, 1) // 根据整型主键查找
	userData := &User{
		Mobile:   "15958165799",
		NickName: "555555",
		Password: "55555",
	}
	db.Model(&User{}).Create(userData)
}
