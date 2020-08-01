package model

import (
	"time"
)

type Comment struct {
	Id          int64      `gorm:"column:id" json:"id"`
	Content     string     `gorm:"column:content" json:"content"`
	CommentTime *time.Time `gorm:"column:comment_time" json:"commentTime"`
	VisitorId   string     `gorm:"column:visitor_id" json:"visitorId"`
	ArticleId   string     `gorm:"column:article_id" json:"articleId"`
}
