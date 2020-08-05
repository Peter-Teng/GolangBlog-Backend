package router

import (
	v1 "MarvelousBlog-Backend/api/v1"
	"github.com/gin-gonic/gin"
)

func LoadVisitorRouters(r *gin.Engine) {
	router := r.Group("/v1/visitor")
	{
		//新增visitor用户
		router.POST("/create", v1.CreateVisitor)

		//获取单个visitor的信息
		router.GET("/detail/:id", v1.GetVisitor)

		//获取多个visitor的信息
		router.GET("/list", v1.GetVisitors)

		//修改visitor信息
		router.PUT("/modify/:id", v1.ModifyVisitor)

		//删除某个visitor
		router.DELETE("/delete/:id", v1.DeleteVisitor)

		//禁用/启用某个visitor
		router.PATCH("/flip/:id", v1.FlipVisitorStatus)

	}
}
