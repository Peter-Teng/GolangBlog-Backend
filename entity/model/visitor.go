package model

import (
	"MarvelousBlog-Backend/common"
	c "MarvelousBlog-Backend/config"
	"net/http"
	"time"
)

type Visitor struct {
	Id            int64      `gorm:"column:id" json:"id" example:"1"`
	Mobile        string     `gorm:"column:mobile;default null" json:"mobile" example:"13xxxxxxxxx"`
	Email         string     `gorm:"column:email;default null" json:"email" example:"xxxxx@xx.com"`
	RegisterTime  *time.Time `gorm:"column:register_time;" json:"registerTime" example:"2020-08-02T21:20:41+08:00"`
	Nickname      string     `gorm:"column:nickname" json:"nickname" example:"PP同学"`
	Password      string     `gorm:"column:password" json:"password" example:"123456"`
	LastLoginTime *time.Time `gorm:"column:last_login_time;" json:"lastLoginTime" example:"2020-08-02T21:20:41+08:00"`
	Status        int8       `gorm:"column:status;default:1" json:"status" example:"1"`
}

type CreateVisitorVO struct {
	Mobile   string `json:"mobile" example:"13xxxxxxxxx"`
	Email    string `json:"email" example:"xxxxx@xx.com"`
	Nickname string `json:"nickname" example:"PP同学"`
	Password string `json:"password" example:"123456"`
}

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
	err := c.Db.Create(&data).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to create Visitor (database), errMsg : ", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusCreated, common.SUCCESS
}

//根据id获取visitor信息
func GetVisitor(id int) (int, int, Visitor) {
	var data Visitor
	c.Db.Where("id = ?", id).First(&data)
	if data.Nickname == "" {
		return http.StatusNotFound, common.VISITOR_NOT_FOUND, data
	}
	return http.StatusOK, common.SUCCESS, data
}
