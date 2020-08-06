package service

import (
	"MarvelousBlog-Backend/common"
	c "MarvelousBlog-Backend/config"
	. "MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
	"github.com/jinzhu/gorm"
	"net/http"
)

//查询visitor用户是否重名
func nickNameUsed(nickname string) bool {
	var visitor Visitor
	c.Db.Select("id").Where("nickname = ?", nickname).First(&visitor)
	//如果查询到的visitor id > 0， 即证明该访客昵称已被占用
	if visitor.Id > 0 {
		//昵称已被占用
		return true
	}
	return false
}

//新增visitor用户
func CreateVisitor(data *Visitor) (int, int) {
	if data.Nickname == "" || data.Password == "" {
		return http.StatusBadRequest, common.EMPTY_VISITOR_INFO
	}
	if nicknameUsed := nickNameUsed(data.Nickname); nicknameUsed {
		return http.StatusForbidden, common.NICKNAME_USED
	}
	var err error
	if data.Password, err = utils.Encrypt(data.Nickname, data.Password); err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to encrypt password", err)
		return http.StatusInternalServerError, common.FAIL
	}
	err = c.Db.Create(&data).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to create Visitor (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusCreated, common.SUCCESS
}

//根据id获取visitor信息
func GetVisitor(id int64) (int, int, Visitor) {
	var data Visitor
	c.Db.Select("id, mobile, email, nickname, register_time, last_login_time, status").Where("id = ?", id).First(&data)
	if data.Nickname == "" {
		return http.StatusNotFound, common.VISITOR_NOT_FOUND, data
	}
	return http.StatusOK, common.SUCCESS, data
}

//获取visitor列表
func ListVisitors(pageSize int, pageNum int) (int, int, []Visitor) {
	var data []Visitor
	err := c.Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&data).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusNotFound, common.VISITOR_NOT_FOUND, nil
		}
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get Visitors (database)", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	return http.StatusOK, common.SUCCESS, data
}

//修改用户信息
func ModifyVisitor(id int64, data *Visitor) (int, int) {
	if data.Nickname == "" {
		return http.StatusBadRequest, common.EMPTY_VISITOR_INFO
	}
	var m = make(map[string]interface{})
	m["mobile"] = data.Mobile
	m["email"] = data.Email
	m["nickname"] = data.Nickname
	if nicknameUsed := nickNameUsed(data.Nickname); nicknameUsed {
		return http.StatusForbidden, common.NICKNAME_USED
	}
	err := c.Db.Model(&data).Where("id = ?", id).Updates(m).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get visitors (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusOK, common.SUCCESS
}

//禁用(启用)visitor
func FlipVisitorStatus(id int64) (int, int) {
	var visitor Visitor
	err := c.Db.Where("id = ?", id).First(&visitor).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to disable visitor (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	var m = make(map[string]interface{})
	m["status"] = 1 - visitor.Status
	err = c.Db.Model(&visitor).Where("id = ?", id).Updates(m).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to disable visitor (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusOK, common.SUCCESS
}

//删除visitor(不推荐)
func DeleteVisitor(id int64) (int, int) {
	var visitor Visitor
	err := c.Db.Where("id = ?", id).Delete(&visitor).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete visitor (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusNoContent, common.SUCCESS
}
