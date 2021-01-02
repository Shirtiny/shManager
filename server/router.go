package server

import (
	"shManager/api"

	"github.com/gin-gonic/gin"
)

// CreateRouter 创建路由
func CreateRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", api.Ping)
	} 

	return router 
}
