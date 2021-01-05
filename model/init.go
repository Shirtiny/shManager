package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
}

// DB 数据库连接单例
var DB *gorm.DB

// ConnectDatabase 连接数据库
func ConnectDatabase() {
	db, err := gorm.Open("mysql", "lab_1591053723:50c712fa6981_#@Aa@tcp(rm-bp1oo27t8762xhlob0o.mysql.rds.aliyuncs.com:3306)/shmysql?charset=utf8&parseTime=True&loc=UTC")
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
	DB.AutoMigrate(&User{})
}
