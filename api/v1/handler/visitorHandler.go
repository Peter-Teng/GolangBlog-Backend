package handler

import (
	"MarvelousBlog-Backend/api/v1/service"
	"MarvelousBlog-Backend/common"
	. "MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity"
	"MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
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
		c.Abort()
		return
	}
	status, code := service.CreateVisitor(&data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Visitor接口
// @Summary 获取单个visitor的信息（以id获取）
// @Description 在URL中输入ID以获取Visitor信息
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "所请求的id参数"
// @Success 200 object model.Visitor "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/visitor/detail/{id} [GET]
func GetVisitor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if (info.IntId != id || info.Role != "visitor") && info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	status, code, data := service.GetVisitor(id)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Visitor接口
// @Summary 获取多个visitor的信息
// @Description 在URL中输入pageNum, pageSize以拉取Visitor列表信息
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param pageSize query int false "请求的页表大小"
// @Param pageNum query int false "请求的offset"
// @Success 200 {array} model.Visitor "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/visitor/list [GET]
func GetVisitors(c *gin.Context) {
	pageNum, err1 := strconv.Atoi(c.DefaultQuery("pageNum", "-1"))
	pageSize, err2 := strconv.Atoi(c.DefaultQuery("pageSize", "-1"))
	if err1 != nil || err2 != nil {
		if err1 != nil {
			utils.ParameterWarnLog(err1)
		} else {
			utils.ParameterWarnLog(err2)
		}
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	status, code, data := service.ListVisitors(pageSize, pageNum)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Visitor接口
// @Summary 修改visitor信息
// @Description 输入新的Visitor信息以更新信息
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "修改的visitor id参数"
// @Param visitor body model.CreateVisitorVO true "访客信息"
// @Success 200 object entity.ResponseObject "修改成功"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 403 object entity.ResponseObject "用户名重复"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/visitor/modify/{id} [PUT]
func ModifyVisitor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if (info.IntId != id || info.Role != "visitor") && info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	var data model.Visitor
	if err := c.ShouldBind(&data); err != nil {
		Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.Abort()
		return
	}
	status, code := service.ModifyVisitor(id, &data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Visitor接口
// @Summary 禁用(启用)某个visitor
// @Description Flip某个visitor的状态，1->0;0->1
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "禁用的visitor id参数"
// @Success 200 object entity.ResponseObject "禁用成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/visitor/flip/{id} [PATCH]
func FlipVisitorStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	status, code := service.FlipVisitorStatus(id)
	if status < 300 {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Visitor接口
// @Summary 删除某个visitor
// @Description 输入Visitor_id以删除visitor
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "删除的visitor id参数"
// @Success 204 "删除成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/visitor/delete/{id}  [DELETE]
func DeleteVisitor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	status, code := service.DeleteVisitor(id)
	if status < 300 {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}
