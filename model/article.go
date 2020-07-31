package model

import (
	"time"
)

type Article struct {
	Id string `gorm:"column:id" json:"id"`
	Title string `gorm:"column:title" json:"title"`
	Summary string `gorm:"column:summary" json:"summary"`
	LabelId string `gorm:"column:label_id" json:"labelId"`
	postTime *time.Time `gorm:"column:post_time" json:"postTime"`
	AuthorId string `gorm:"column:author_id" json:"authorId"`
	VisitCount string `gorm:"column:visit_count" json:"visitCount"`
	lastModifyTime *time.Time `gorm:"column:last_modify_count" json:"lastModifyTime"`
}
