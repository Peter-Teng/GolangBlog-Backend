package service

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
	"net/http"
)

func VisitorLogin(nickname string, password string) (status int, code int, token string) {
	var data model.Visitor
	var err error
	if err = config.Db.Select("id, password").Where("nickname = ?", nickname).First(&data).Error; err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "login finding visitor failed", err)
		return http.StatusInternalServerError, common.FAIL, ""
	}
	if data.Id == 0 {
		return http.StatusNotFound, common.VISITOR_NOT_FOUND, ""
	}
	pwd, err := utils.Encrypt(nickname, password)
	if err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "encrypt password failed", err)
		return http.StatusInternalServerError, common.FAIL, ""
	}
	if pwd != data.Password {
		return http.StatusForbidden, common.NAME_OR_PASSWORD_ERROR, ""
	}
	//传入身份为visitor
	token, err = utils.GenerateToken(data.Nickname, "visitor", data.Id)
	if err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "generate token failed", err)
		return http.StatusInternalServerError, common.FAIL, ""
	}
	return http.StatusOK, common.SUCCESS, token
}
