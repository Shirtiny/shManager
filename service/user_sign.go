package service

import (
	"fmt"
	"shManager/auth"
	"shManager/model"
	"shManager/serializer"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// UserSignService 用户注册、登陆服务的接口
type UserSignService interface {
	SignUp(name string, password string, nickname string) (serializer.User, error)
	SignIn(name string, password string) (string, error)
}

// userSignServiceImpl 用户注册、登陆的实现类
type userSignServiceImpl struct {
	passWordCost int
}

// NewUserSignService 返回一个接口实现的实例
func NewUserSignService() UserSignService {
	return &userSignServiceImpl{
		// 密码加密等级
		passWordCost: 12,
	}
}

//用户注册
func (service *userSignServiceImpl) SignUp(name string, password string, nickname string) (serializer.User, error) {
	// 加密用户密码
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), service.passWordCost)
	if err != nil {
		return serializer.User{}, err
	}
	user := model.User{
		// 将用户名重置为小写
		Name: strings.ToLower(name),
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

// 用户登陆
func (service *userSignServiceImpl) SignIn(name, password string) (string, error) {
	// 根据用户名查询用户
	user, err := model.UserGet(model.User{
		Name: name,
	})
	if err != nil {
		return "", err
	}
	// 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}
	// 生成jwt
	jwt, err := auth.CreateJwt(serializer.BuildUser(user))
	if err != nil {
		return "", err
	}
	return jwt, nil
}
