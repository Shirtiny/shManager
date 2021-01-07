package server

import (
	"shManager/api"

	"github.com/gin-gonic/gin"
)

// CreateRouter 创建路由
func CreateRouter(prefixPath string) *gin.Engine {
	router := gin.Default()

	v1 := router.Group(prefixPath + "/v1")
	{
		v1.GET("/ping", api.Ping)
		v1.GET("/user", api.User)
	}

	return router
}
