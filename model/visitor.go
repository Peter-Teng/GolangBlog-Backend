package model

import "time"

type Visitor struct {
	Id            string     `gorm:"column:id" json:"id"`
	Mobile        string     `gorm:"column:mobile" json:"mobile"`
	Email         string     `gorm:"column:email" json:"email"`
	RegisterTime  *time.Time `gorm:"column:register_time" json:"registerTime"`
	Nickname      string     `gorm:"column:nickname" json:"nickname"`
	Password      string     `gorm:"column:password" json:"password"`
	LastLoginTime *time.Time `gorm:"column:last_login_time;" json:"lastLoginTime"`
	Status        int8       `gorm:"column:status;default 1" json:"status"`
}
