package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 数据库连接单例
var DB *gorm.DB

// ConnectDatabase 连接数据库
func ConnectDatabase(dsn string) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败", err)
		panic(err)
	}
	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	// 自动迁移模式
	DB.AutoMigrate(&Key{})
	DB.AutoMigrate(&User{})
}
