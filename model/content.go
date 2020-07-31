package model

type Content struct {
	Id string `gorm:"column:id" json:"id"`
	ArticleId string `gorm:"column:article_id" json:"articleId"`
	Content string `gorm:"column:content" json:"content"`
}