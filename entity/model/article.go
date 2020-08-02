package model

import (
	"time"
)

type Article struct {
	Id             int64      `gorm:"column:id" json:"id"`
	Title          string     `gorm:"column:title" json:"title"`
	Summary        string     `gorm:"column:summary" json:"summary"`
	LabelId        int64      `gorm:"column:label_id" json:"labelId"`
	PostTime       *time.Time `gorm:"column:post_time;" json:"postTime"`
	AuthorId       int64      `gorm:"column:author_id" json:"authorId"`
	VisitCount     int64      `gorm:"column:visit_count" json:"visitCount"`
	LastModifyTime *time.Time `gorm:"column:last_modify_time" json:"lastModifyTime"`
}
