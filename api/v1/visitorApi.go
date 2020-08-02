package v1

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

//新增visitor用户
func CreateVisitor(c *gin.Context) {
	var data model.Visitor
	_ = c.ShouldBind(&data)
	status, code := model.CreateVisitor(&data)
	c.JSON(status, gin.H{
		"code":    code,
		"message": common.Message[code],
	})
}

//获取单个visitor的信息（以id获取）
func GetVisitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, code, data := model.GetVisitor(id)
	c.JSON(status, gin.H{
		"code":    code,
		"message": common.Message[code],
		"data":    data,
	})
}

//获取多个visitor的信息
func GetVisitors(c *gin.Context) {

}

//修改visitor信息
func ModifyVisitor(c *gin.Context) {

}

//禁用某个visitor
func DisableVisitor(c *gin.Context) {

}
