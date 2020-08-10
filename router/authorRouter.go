package router

import (
	"MarvelousBlog-Backend/api/v1/handler"
	"github.com/gin-gonic/gin"
)

func LoadAuthorRouters(r *gin.Engine) {
	router := r.Group("/v1/author")
	{
		//author登录接口
		router.POST("/login", handler.AuthorLogin)

	}
}
