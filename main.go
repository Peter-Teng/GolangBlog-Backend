package main

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/middleware"
	r "MarvelousBlog-Backend/router"
	"github.com/gin-gonic/gin"
)

// @title PP同学个人博客接口文档
// @version 0.1
// @description MarvelousBlog-Backend Swagger接口文档
// @contact.name PP同学
// @contact.email 710955321@qq.com
// @host localhost:8600
func main() {
	gin.SetMode(config.AppMode)
	router := gin.New()

	//添加日志中间件
	router.Use(middleware.LoggingMiddleware())

	//添加错误恢复中间件
	router.Use(gin.Recovery())

	//加载各类router
	r.LoadVisitorRouters(router)

	//启动swagger
	config.InitSwaggerRouter(router)

	//启动服务
	err := router.Run(config.ServerPort)
	if err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Service run failed! errMsg = ", err)
		closeResources()
	}

	//关闭资源
	closeResources()
}

func closeResources() {
	config.Log.Info("Service Shutdown!")
	if err := config.RedisPool.Close(); err != nil {
		config.Log.Error(common.SYSTEM_ERROR_LOG, "Redis pool closed error! errMsg = ", err)
	}
}
