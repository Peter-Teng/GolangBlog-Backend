package router

import (
	"MarvelousBlog-Backend/api/v1/handler"
	"MarvelousBlog-Backend/middleware"
	"github.com/gin-gonic/gin"
)

func LoadLabelRouters(r *gin.Engine) {
	router := r.Group("/v1/label")
	{
		//新增label
		router.POST("/create", middleware.Auth(), handler.CreateLabel)

		//获取全部label
		router.GET("/list", handler.GetAllLabels)

		//修改label信息
		router.PUT("/modify/:id", middleware.Auth(), handler.ModifyLabel)

		//删除某个label
		router.DELETE("/delete/:id", middleware.Auth(), handler.DeleteLabel)

	}
}
