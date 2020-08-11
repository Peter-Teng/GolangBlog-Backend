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

// @Tags Comment接口
// @Summary 新增评论
// @Description 为某篇文章添加评论
// @Accept  json
// @Produce json
// @Param label body model.Comment true "评论内容"
// @Success 201 object entity.ResponseObject "评论成功"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/comment/create [POST]
func MakeComment(c *gin.Context) {
	var data model.Comment
	if err := c.ShouldBind(&data); err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	if msg, ok := utils.Validate(&data); !ok {
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, msg))
		c.Abort()
		return
	}
	if data.Mobile == "" {
		data.Mobile = " Anonymous "
	}
	status, code := service.MakeComment(&data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Comment接口
// @Summary 分页获取某篇文章下的评论
// @Description 获取某篇文章文章的20个评论
// @Accept  json
// @Produce json
// @Param pageNum query int false "请求的页码数"
// @Param articleId path int true "评论的article的id"
// @Success 200 {array} model.Comment "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/comment/list/{articleId} [GET]
func ListComment(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	id, err := strconv.ParseInt(c.Param("articleId"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	status, code, data := service.ListComment(id, pageNum)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Comment接口
// @Summary 删除某条评论
// @Description 输入评论id以删除某条评论
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "删除的评论的id"
// @Success 204 "删除成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/comment/delete/{id}  [DELETE]
func DeleteComment(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists || claims.(*utils.BlogClaims).Role != "superAuthor" {
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
	status, code := service.DeleteComment(id)
	if status < 300 {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}
