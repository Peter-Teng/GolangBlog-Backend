package model

type Author struct {
	Id        string `gorm:"column:id" json:"id"`
	LoginName string `gorm:"column:login_name" json:"loginName"`
	Password  string `gorm:"column:password" json:"password"`
	Icon      string `gorm:"column:icon" json:"icon"`
}
