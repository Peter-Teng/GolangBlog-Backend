package model

type Author struct {
	Id       int64  `gorm:"column:id" json:"id"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Password string `gorm:"column:password" json:"password"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Role     int8   `gorm:"column:role" json:"role"`
}
