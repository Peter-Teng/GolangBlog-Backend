package router

import (
	"MarvelousBlog-Backend/api/v1/handler"
	"MarvelousBlog-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func LoadArticleRouter(r *gin.Engine) {
	router := r.Group("/v1/article")
	{
		//新增文章
		router.POST("/create", middleware.Auth(), handler.CreateArticle)

		//获取文章
		router.GET("/detail/:id", handler.GetArticle)

		//获取文章列表
		router.GET("/list", handler.GetArticles)

		//编辑文章
		router.PUT("/modify/:id", middleware.Auth(), handler.ModifyArticle)

		//删除某篇文章
		router.DELETE("/delete/:id", middleware.Auth(), handler.DeleteArticle)

		//获取某一个Label下的所有文章
		router.GET("/onLabel/:labelId", handler.GetArticlesByLabel)

		//管理员获取文章列表（可以获取到status为0的文章）
		router.GET("/superAuthor/list", middleware.Auth(), handler.ManageArticles)

		//重新恢复文章（status 0 -> 1）
		router.PATCH("/enable/:id", middleware.Auth(), handler.EnableArticle)
	}
}
