package model

import (
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

type VisitorLoginVo struct {
	Nickname string `json:"nickname" example:"PP同学"`
	Password string `json:"password" example:"123456"`
}
