package main

import (
	"MarvelousBlog-Backend/config"
	r "MarvelousBlog-Backend/router"
	"fmt"
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
	router := gin.Default()

	//加载各类router
	r.LoadVisitorRouters(router)

	//启动swagger
	config.InitSwaggerRouter(router)

	//启动服务
	err := router.Run(config.ServerPort)
	if err != nil {
		fmt.Println("Service run failed!")
	}
}
