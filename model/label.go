package model

type Label struct {
	Id string `gorm:"column:id" json:"id"`
	LabelName string `gorm:"column:label_name" json:"labeName"`
	Description string `gorm:"column:description" json:"description"`
}
