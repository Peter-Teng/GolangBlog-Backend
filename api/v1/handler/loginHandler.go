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
)

// @Tags 登录接口
// @Summary author登录
// @Description 输入用户名密码以登录
// @Accept  json
// @Produce json
// @Param visitor body model.LoginVo true "作者登录信息"
// @Success 200 object entity.ResponseObject "登录成功"
// @Failure 403 object entity.ResponseObject "用户名或密码错误"
// @Failure 404 object entity.ResponseObject "未找到该用户"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/author/login  [POST]
func AuthorLogin(c *gin.Context) {
	var data model.Author
	if err := c.ShouldBind(&data); err != nil {
		Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, common.Message[common.PARAMETER_BAD_REQUEST]))
		c.Abort()
		return
	}
	if msg, ok := utils.Validate(&data); !ok {
		c.JSON(http.StatusBadRequest, entity.NewResponseObject(common.PARAMETER_BAD_REQUEST, msg))
		c.Abort()
		return
	}
	status, code, token := service.AuthorLogin(data.Nickname, data.Password)
	if status < 300 {
		c.JSON(status, gin.H{
			"code":    code,
			"message": common.Message[code],
			"token":   token,
		})
	} else {
		c.JSON(status, entity.NewResponseObject(code, common.Message[code]))
	}
}
