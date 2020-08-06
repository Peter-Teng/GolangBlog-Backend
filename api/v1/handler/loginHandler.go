package handler

import (
	"MarvelousBlog-Backend/api/v1/service"
	"MarvelousBlog-Backend/common"
	. "MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity"
	"MarvelousBlog-Backend/entity/model"
	"github.com/gin-gonic/gin"
)

// @Tags 登录接口
// @Summary visitor登录
// @Description 输入用户名密码以登录
// @Accept  json
// @Produce json
// @Param visitor body model.VisitorLoginVo true "访客登录信息（只需要填写nickname、password）"
// @Success 200 object entity.ResponseObject "登录成功"
// @Failure 403 object entity.ResponseObject "用户名或密码错误"
// @Failure 404 object entity.ResponseObject "未找到该用户"
// @Failure 500 object entity.ResponseObject "服务器错误"
// @Router /v1/visitor/login  [POST]
func VisitorLogin(c *gin.Context) {
	var data model.Visitor
	if err := c.ShouldBind(&data); err != nil {
		Log.Errorf(common.SYSTEM_ERROR_LOG, "Json bind error!", err)
		c.Abort()
		return
	}
	status, code, token := service.VisitorLogin(data.Nickname, data.Password)
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
