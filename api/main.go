package api

import (
	"shManager/serializer"

	"github.com/gin-gonic/gin"
)

// Ping ping
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong333",
	})
}
