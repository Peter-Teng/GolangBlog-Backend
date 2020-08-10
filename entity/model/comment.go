package model

import (
	"time"
)

type Comment struct {
	Id          int64      `gorm:"column:id" json:"id" swaggerignore:"true"`
	Mobile      string     `gorm:"mobile" json:"mobile" example:"137xxxxxxxx"`
	Email       string     `gorm:"email" json:"email" example:"1846156416@qq.com"`
	Content     string     `gorm:"column:content" json:"content" example:"写得不错！"`
	CommentTime *time.Time `gorm:"column:comment_time" json:"commentTime" swaggerignore:"true"`
	ArticleId   int64      `gorm:"column:article_id" json:"articleId" example:"1"`
}
