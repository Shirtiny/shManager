package service

import (
	"fmt"
	"shManager/model"
	"shManager/serializer"
)

// UserSignUp 用户注册
func UserSignUp(name string, password string, nickname string) (serializer.User, error) {
	user := model.User{
		Name:     name,
		Password: password,
		Nickname: nickname,
	}
	fmt.Println("开始注册用户：", user)
	err := model.UserAdd(user)
	if err != nil {
		return serializer.User{}, err
	}
	return serializer.BuildUser(user), nil
}
