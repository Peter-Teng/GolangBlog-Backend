package router

import (
	"MarvelousBlog-Backend/api/v1/handler"
	"MarvelousBlog-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func LoadVisitorRouters(r *gin.Engine) {
	router := r.Group("/v1/visitor")
	{
		//新增visitor用户
		router.POST("/create", handler.CreateVisitor)

		//获取单个visitor的信息
		router.GET("/detail/:id", middleware.Auth(), handler.GetVisitor)

		//获取多个visitor的信息
		router.GET("/list", middleware.Auth(), handler.GetVisitors)

		//修改visitor信息
		router.PUT("/modify/:id", middleware.Auth(), handler.ModifyVisitor)

		//删除某个visitor
		router.DELETE("/delete/:id", middleware.Auth(), handler.DeleteVisitor)

		//禁用/启用某个visitor
		router.PATCH("/flip/:id", middleware.Auth(), handler.FlipVisitorStatus)

		//visitor登录接口
		router.POST("/login", handler.VisitorLogin)
	}
}
