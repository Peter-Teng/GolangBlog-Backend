package router

import (
	"MarvelousBlog-Backend/api/v1/handler"
	"MarvelousBlog-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func LoadCommentRouter(r *gin.Engine) {
	router := r.Group("/v1/comment")
	{
		//新增评论
		router.POST("/create", handler.MakeComment)

		//获取某篇文章的全部评论
		router.GET("/list/:articleId", handler.ListComment)

		//删除某个评论
		router.DELETE("/delete/:id", middleware.Auth(), handler.DeleteComment)
	}
}
