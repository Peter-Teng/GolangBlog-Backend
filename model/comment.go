package model

import (
	"time"
)

type Comment struct {
	Id string `gorm:"column:id" json:"id"`
	content string `gorm:"column:content" json:"content"`
	commentTime *time.Time `gorm:"column:comment_time" json:"comment_time"`
	VisitorId string `gorm:"column:visitor_id" json:"visitor_id"`
	ArticleId string `gorm:"column:article_id" json:"article_id"`
}
