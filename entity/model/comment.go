package model

import (
	"time"
)

type Comment struct {
	Id          int64      `gorm:"column:id" json:"id"`
	Mobile      string     `gorm:"mobile" json:"mobile"`
	Email       string     `gorm:"email" json:"email"`
	Content     string     `gorm:"column:content" json:"content"`
	CommentTime *time.Time `gorm:"column:comment_time" json:"commentTime"`
	ArticleId   string     `gorm:"column:article_id" json:"articleId"`
	Status      int8       `gorm:"status" json:"status"`
}
