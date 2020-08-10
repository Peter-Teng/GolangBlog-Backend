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

// @Tags Article接口
// @Summary 新增文章
// @Description 创建一篇新的文章
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param label body model.Article true "文章内容"
// @Success 201 object entity.ResponseObject "新增文章成功"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/article/create [POST]
func CreateArticle(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists || claims.(*utils.BlogClaims).Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}

	var data model.Article
	if err := c.ShouldBind(&data); err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.Abort()
		return
	}
	status, code := service.CreateArticle(&data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Article接口
// @Summary 获取某篇文章
// @Description 获取某篇文章的内容
// @Accept  json
// @Produce json
// @Param id path int true "所请求的文章id"
// @Success 200 {array} model.Article "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/article/detail/{id} [GET]
func GetArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	status, code, data := service.GetArticle(id)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Article接口
// @Summary 分页获取全部文章(每页10篇）
// @Description 获取某篇文章的内容
// @Accept  json
// @Produce json
// @Param pageNum query int false "请求的页码数"
// @Success 200 {array} model.Article "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/article/list [GET]
func GetArticles(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	status, code, data := service.GetArticles(pageNum)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Article接口
// @Summary 修改文章内容
// @Description 输入新的文章内容以更新文章
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "修改的文章的id"
// @Param visitor body model.Article true "修改的文章信息"
// @Success 200 object entity.ResponseObject "修改成功"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/article/modify/{id} [PUT]
func ModifyArticle(c *gin.Context) {
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
	var data model.Article
	if err := c.ShouldBind(&data); err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.Abort()
		return
	}
	status, code := service.ModifyArticle(id, &data)
	c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
}

// @Tags Article接口
// @Summary 删除某篇文章
// @Description 输入文章id以删除文章（软删除，并未在数据库中实际删除）
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "删除的article的id"
// @Success 204 "删除成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 402 object entity.ResponseObject "用户未授权"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/article/delete/{id}  [DELETE]
func DeleteArticle(c *gin.Context) {
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
	status, code := service.DeleteArticle(id)
	if status < 300 {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Article接口
// @Summary 获取某一标签下的全部文章
// @Description 获取某一标签下的全部文章
// @Accept  json
// @Produce json
// @Param labelId path int true "请求的标签的id"
// @Success 200 {array} model.Article "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/article/onLabel/{labelId} [GET]
func GetArticlesByLabel(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("labelId"), 10, 64)
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	status, code, data := service.GetArticlesByLabel(id)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Article接口
// @Summary 管理员获取全部文章
// @Description 管理员获取全部文章（包括status = 0的）
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param pageNum query int false "请求的页码数"
// @Success 200 {array} model.Article "查询成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/article/superAuthor/list [GET]
func ManageArticles(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists || claims.(*utils.BlogClaims).Role != "superAuthor" {
		c.JSON(http.StatusUnauthorized, entity.NewResponseObject(common.UNAUTHORIZED, common.Message[common.UNAUTHORIZED]))
		c.Abort()
		return
	}
	pageNum, err := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	if err != nil {
		utils.ParameterWarnLog(err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	status, code, data := service.ManageArticles(pageNum)
	if status < 300 {
		c.JSON(status, data)
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}

// @Tags Article接口
// @Summary 管理员Enable文章
// @Description 管理员重新启用某篇文章（status 0 -> 1）
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param id path int true "enable的文章id"
// @Success 200 object entity.ResponseObject "Enable成功"
// @Failure 400 object entity.ResponseObject "输入参数有误"
// @Failure 404 object entity.ResponseObject "未找到资源"
// @Router /v1/article/enable/{id} [PATCH]
func EnableArticle(c *gin.Context) {
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
	status, code := service.EnableArticle(id)
	if status < 300 {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}
