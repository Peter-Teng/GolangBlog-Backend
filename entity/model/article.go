package model

import (
	"time"
)

type Article struct {
	Id             int64      `gorm:"column:id" json:"id"`
	Title          string     `gorm:"column:title" json:"title"`
	Abstract       string     `gorm:"column:abstract" json:"summary"`
	Content        string     `gorm:"column:content" json:"content"`
	LabelId        int64      `gorm:"column:label_id" json:"labelId"`
	PostTime       *time.Time `gorm:"column:post_time;" json:"postTime"`
	AuthorId       int64      `gorm:"column:author_id" json:"authorId"`
	VisitCount     int64      `gorm:"column:visit_count" json:"visitCount"`
	LastModifyTime *time.Time `gorm:"column:last_modify_time" json:"lastModifyTime"`
	Status         int8       `gorm:"column:status" json:"status"`
	Label          Label      `json:"label" swaggerignore:"true"`
	Comments       []Comment  `json:"comments" swaggerignore:"true"`
}

type ArticleVO struct {
	Id               int64
	Title            string
	Abstract         string
	Post_time        *time.Time
	Author_id        int64
	Visit_count      int64
	Last_modify_time *time.Time
	Label_name       string
	Description      string
	status           int8
}
