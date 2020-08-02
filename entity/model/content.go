package model

type Content struct {
	Id        int64  `gorm:"column:id" json:"id"`
	ArticleId int64  `gorm:"column:article_id" json:"articleId"`
	Content   string `gorm:"column:content" json:"content"`
}
