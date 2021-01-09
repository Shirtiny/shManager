package api

import (
	"fmt"
	"shManager/model"
	"shManager/serializer"
	"shManager/service"

	"github.com/gin-gonic/gin"
)

// GetUser 获取用户
func GetUser(c *gin.Context) {
	var user model.User

	model.DB.First(&user)
	fmt.Println(user)
	c.JSON(200, user)
}

// UserSignUp 用户注册
func UserSignUp(c *gin.Context) {
	user := model.User{}
	c.ShouldBind(&user)
	createdUser, err := service.UserSignUp(user.Name, user.Password, user.Nickname)
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
