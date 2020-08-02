package v1

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/entity"
	"MarvelousBlog-Backend/entity/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Tags Visitor接口
// @Summary 新增visitor用户
// @Description 输入信息来创建一个visitor
// @Accept  json
// @Produce json
// @Param mobile body string false "访客电话号码"
// @Param email body string false "访客邮箱"
// @Param nickname body string true "访客昵称"
// @Param password body string true "访客密码"
// @Success 201 object entity.ResponseBody
// @Router /v1/visitor/create [POST]
func CreateVisitor(c *gin.Context) {
	var data model.Visitor
	_ = c.ShouldBind(&data)
	status, code := model.CreateVisitor(&data)
	c.JSON(status, entity.NewResponseBody(code, common.Message[code]))
}

// @Tags Visitor接口
// @Summary 获取单个visitor的信息（以id获取）
// @Description 在URL中输入ID以获取Visitor信息
// @Accept  json
// @Produce json
// @Param id path int true "所请求的id参数"
// @Success 200 object model.Visitor
// @Router /v1/visitor/info/{id} [GET]
func GetVisitor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, code, data := model.GetVisitor(id)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseBody(code, common.Message[code]))
	}
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
