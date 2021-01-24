package api

import (
	"fmt"
	"shManager/auth"
	"shManager/model"
	"shManager/serializer"
	"shManager/service"

	"github.com/gin-gonic/gin"
)

var userSignService service.UserSignService

func init() {
	fmt.Println("实例化用户注册登陆接口")
	userSignService = service.NewUserSignService()
}

// GetUser 获取用户
func GetUser(c *gin.Context) {
	token := c.Request.Header.Get("token")
	claims, err := auth.ParseJwt(token)
	if err != nil {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "token无效或过期",
			Data: nil,
		})
		return
	}
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "ok",
		Data: claims,
	})
}

// UserSignUp 用户注册
func UserSignUp(c *gin.Context) {
	user := model.User{}
	c.ShouldBind(&user)

	createdUser, err := userSignService.SignUp(user.Name, user.Password, user.Nickname)
	if err != nil {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(200, serializer.Response{
			Code: 0,
			Msg:  "ok",
			Data: createdUser,
		})
	}
}

// UserLogin 用户登陆
func UserLogin(c *gin.Context) {
	user := model.User{}
	c.ShouldBind(&user)
	token, err := userSignService.SignIn(user.Name, user.Password)
	if err != nil {
		c.JSON(200, serializer.Response{
			Code: 1,
			Msg:  "用户名或密码错误",
			Data: nil,
		})
		return
	}
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "ok",
		Data: token,
	})
}
