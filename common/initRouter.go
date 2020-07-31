package common

import (
	"MarvelousBlog-Backend/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 添加 Get 请求路由
	router.GET("/", api.Hello)
	return router
}
