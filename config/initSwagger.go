package config

import (
	"MarvelousBlog-Backend/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//初始化swagger文档
func InitSwaggerRouter(r *gin.Engine) {
	docs.SwaggerInfo.Host = ServerHost + ServerPort
	url := ginSwagger.URL("http://" + ServerHost + ServerPort + "/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
