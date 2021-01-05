package api

import (
	"fmt"
	"shManager/model"
	"shManager/serializer"

	"github.com/gin-gonic/gin"
)

// Ping ping
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
}

// User 获取用户
func User(c *gin.Context) {
	var user model.User

	model.DB.First(&user)
	fmt.Println(user)
	c.JSON(200, user)

}
