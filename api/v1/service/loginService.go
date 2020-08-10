package service

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
	"net/http"
)

//作者登录逻辑
func AuthorLogin(nickname string, password string) (status int, code int, token string) {
	var data model.Author
	var err error
	if err = config.Db.Select("id, password, role").Where("nickname = ?", nickname).FirstOrInit(&data).Error; err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "login finding author failed", err)
		return http.StatusInternalServerError, common.FAIL, ""
	}
	if data.Id == 0 {
		return http.StatusNotFound, common.USER_NOT_FOUND, ""
	}
	pwd, err := utils.Encrypt(nickname, password)
	if err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "encrypt password failed", err)
		return http.StatusInternalServerError, common.FAIL, ""
	}
	if pwd != data.Password {
		return http.StatusForbidden, common.NAME_OR_PASSWORD_ERROR, ""
	}
	//传入身份信息
	if data.Role == 1 {
		token, err = utils.GenerateToken(data.Nickname, "superAuthor", data.Id)
		if err != nil {
			config.Log.Errorf(common.SYSTEM_ERROR_LOG, "generate token failed", err)
			return http.StatusInternalServerError, common.FAIL, ""
		}
	} else {
		token, err = utils.GenerateToken(data.Nickname, "author", data.Id)
		if err != nil {
			config.Log.Errorf(common.SYSTEM_ERROR_LOG, "generate token failed", err)
			return http.StatusInternalServerError, common.FAIL, ""
		}
	}
	UpdateLastLoginTime(data.Id)
	return http.StatusOK, common.SUCCESS, token
}

//更新登录时间
func UpdateLastLoginTime(id int64) {
	var m = make(map[string]interface{})
	m["last_login_time"] = nil
	config.Db.Model(&model.Author{}).Where("id = ?", id).Update(m)
}
