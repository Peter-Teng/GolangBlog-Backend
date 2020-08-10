package model

type Label struct {
	Id          int64     `gorm:"column:id" json:"id" example:"1"`
	LabelName   string    `gorm:"column:label_name" json:"labeName" example:"Gin框架"`
	Description string    `gorm:"column:description" json:"description" example:"如何使用Gin框架搭建一个web服务器"`
	Articles    []Article `gorm:"foreignKey:AuthorId" json:"articles" swaggerignore:"true"`
}
