package service

import (
	"fmt"
	"shManager/model"
	"shManager/serializer"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	// PassWordCost 密码加密难度
	PassWordCost = 12
)

// UserSignUp 用户注册
func UserSignUp(name string, password string, nickname string) (serializer.User, error) {
	// 加密用户密码
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return serializer.User{}, err
	}
	user := model.User{
		// 将用户名重置为小写
		Name:     strings.ToLower(name),
		// 加密后的密码
		Password: string(pwdBytes),
		Nickname: nickname,
	}
	fmt.Println("开始创建用户：", user)

	dbErr := model.UserAdd(user)
	if dbErr != nil {
		return serializer.User{}, dbErr
	}
	return serializer.BuildUser(user), nil
}
