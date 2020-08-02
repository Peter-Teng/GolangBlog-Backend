package model

import (
	"MarvelousBlog-Backend/common"
	c "MarvelousBlog-Backend/config"
	"fmt"
	"net/http"
	"time"
)

type Visitor struct {
	Id            int64      `gorm:"column:id" json:"id"`
	Mobile        string     `gorm:"column:mobile;default null" json:"mobile"`
	Email         string     `gorm:"column:email;default null" json:"email"`
	RegisterTime  *time.Time `gorm:"column:register_time;" json:"registerTime"`
	Nickname      string     `gorm:"column:nickname" json:"nickname"`
	Password      string     `gorm:"column:password" json:"password"`
	LastLoginTime *time.Time `gorm:"column:last_login_time;" json:"lastLoginTime"`
	Status        int8       `gorm:"column:status;default:1" json:"status"`
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
	if nicknameUsed := nickNameUsed(data.Nickname); nicknameUsed {
		return http.StatusForbidden, common.NICKNAME_USED
	}
	err := c.Db.Create(&data).Error
	if err != nil {
		fmt.Println(err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusCreated, common.SUCCESS
}

func GetVisitor(id int) (int, int, Visitor) {
	var data Visitor
	c.Db.Where("id = ?", id).First(&data)
	if data.Nickname == "" {
		return http.StatusNotFound, common.VISITOR_NOT_FOUND, data
	}
	return http.StatusOK, common.SUCCESS, data
}
