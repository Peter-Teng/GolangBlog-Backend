package model

type Author struct {
	Id        int64  `gorm:"column:id" json:"id"`
	LoginName string `gorm:"column:login_name" json:"loginName"`
	Password  string `gorm:"column:password" json:"password"`
	Avatar    string `gorm:"column:avatar" json:"avatar"`
	Role      int8   `gorm:"column:role" json:"role"`
}
