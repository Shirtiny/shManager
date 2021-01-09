package model

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name     string `gorm:"not null;unique;"`
	Password string `gorm:"not null"`
	Nickname string
	Avatar   string `gorm:"default:'默认头像'"`
	Email    string
}

// UserAdd 添加用户
func UserAdd(user User) error {
	if err := DB.Create(&user).Error; err != nil {
		return errors.New("插入用户失败：" + err.Error())
	}
	return nil
}
