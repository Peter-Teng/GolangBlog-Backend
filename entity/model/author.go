package model

import "time"

type Author struct {
	Id            int64      `gorm:"column:id" json:"id"`
	Nickname      string     `gorm:"column:nickname" json:"nickname"`
	Password      string     `gorm:"column:password" json:"password"`
	Avatar        string     `gorm:"column:avatar" json:"avatar"`
	LastLoginTime *time.Time `gorm:"column:last_login_time;" json:"lastLoginTime" example:"2020-08-02T21:20:41+08:00"`
	RegisterTime  *time.Time `gorm:"column:register_time;" json:"registerTime" example:"2020-08-02T21:20:41+08:00"`
	Role          int8       `gorm:"column:role" json:"role"`
}

type LoginVo struct {
	Nickname string `json:"nickname" example:"PP同学"`
	Password string `json:"password" example:"123456"`
}
