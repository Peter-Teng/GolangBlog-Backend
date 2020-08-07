package handler

import (
	"MarvelousBlog-Backend/api/v1/service"
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity"
	"MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Tags Label接口
// @Summary 新增Label标签
// @Description 输入信息来创建一个标签
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param label body model.Label true "新建的标签信息"
// @Success 201 object entity.ResponseObject "新增标签成功"
// @Failure 403 object entity.ResponseObject "标签重复"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/label/create [POST]
func CreateLabel(c *gin.Context) {
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}

	var data model.Label
	if err := c.ShouldBind(&data); err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.Abort()
		return
	}
	status, code := service.CreateLabel(&data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Label接口
// @Summary 获取全部label
// @Description 获取全部label
// @Accept  json
// @Produce json
// @Success 200 {array} model.Label "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/label/list [GET]
func GetAllLabels(c *gin.Context) {
	status, code, data := service.GetAllLabels()
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Label接口
// @Summary 修改label
// @Description 输入修改的label信息以更新信息
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "修改的label id参数"
// @Param label body model.Label true "修改的标签信息"
// @Success 200 object entity.ResponseObject "修改成功"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 403 object entity.ResponseObject "Label重复"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/label/modify/{id} [PUT]
func ModifyLabel(c *gin.Context) {
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
	var data model.Label
	if err := c.ShouldBind(&data); err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.Abort()
		return
	}
	status, code := service.ModifyLabel(id, &data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Label接口
// @Summary 删除某个label
// @Description 输入label_id以删除标签
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "删除的label的id"
// @Success 204 "删除成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/label/delete/{id}  [DELETE]
func DeleteLabel(c *gin.Context) {
	claims, _ := c.Get("claims")
	info := claims.(*utils.BlogClaims)
	if info.Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	status, code := service.DeleteLabel(id)
	if status < 300 {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}
