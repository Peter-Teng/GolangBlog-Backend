package v1

import (
	"MarvelousBlog-Backend/common"
	. "MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity"
	"MarvelousBlog-Backend/entity/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags Visitor接口
// @Summary 新增visitor用户
// @Description 输入信息来创建一个visitor
// @Accept  json
// @Produce json
// @Param visitor body model.CreateVisitorVO true "注册访客信息"
// @Success 201 object entity.ResponseObject "注册成功"
// @Failure 403 object entity.ResponseObject "用户名重复"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/visitor/create [POST]
func CreateVisitor(c *gin.Context) {
	var data model.Visitor
	if err := c.ShouldBind(&data); err != nil {
		Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
	}
	status, code := model.CreateVisitor(&data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Visitor接口
// @Summary 获取单个visitor的信息（以id获取）
// @Description 在URL中输入ID以获取Visitor信息
// @Accept  json
// @Produce json
// @Param id path int true "所请求的id参数"
// @Success 200 object model.Visitor "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/visitor/detail/{id} [GET]
func GetVisitor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		Log.Warnf(common.SYSTEM_ERROR_LOG, "输入参数错误", err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		return
	}
	status, code, data := model.GetVisitor(id)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Visitor接口
// @Summary 获取多个visitor的信息
// @Description 在URL中输入pageNum, pageSize以拉取Visitor列表信息
// @Accept  json
// @Produce json
// @Param pageSize query int false "请求的页表大小"
// @Param pageNum query int false "请求的offset"
// @Success 200 {array} model.Visitor "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/visitor/list [GET]
func GetVisitors(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	if err != nil {
		Log.Warnf(common.SYSTEM_ERROR_LOG, "输入参数错误", err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		return
	}
	status, code, data := model.ListVisitors(pageSize, pageNum)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

//修改visitor信息
func ModifyVisitor(c *gin.Context) {

}

//禁用某个visitor
func DisableVisitor(c *gin.Context) {

}

func DeleteVisitor(c *gin.Context) {

}
