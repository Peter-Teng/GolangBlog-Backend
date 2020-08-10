package service

import (
	"MarvelousBlog-Backend/common"
	c "MarvelousBlog-Backend/config"
	. "MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

const (
	QUERY_ARTICLE_BY_ID_KEY     = "ARTICLE_%d"
	QUERY_ARTICLES_BY_LABEL_KEY = "ARTICLES_ON_LABEL_%d"
)

func CreateArticle(data *Article) (int, int) {
	if err := c.Db.Create(&data).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to create Article (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusCreated, common.SUCCESS
}

func GetArticle(id int64) (int, int, Article) {
	var data Article
	article, err := utils.RedisGet(fmt.Sprintf(QUERY_ARTICLE_BY_ID_KEY, id))
	if err == nil {
		err = c.Json.Unmarshal(utils.Str2bytes(article), &data)
		if err == nil {
			return http.StatusOK, common.SUCCESS, data
		}
	}
	c.Db.Preload("Label").Where("id = ?", id).First(&data)
	if data.Title == "" || data.Status <= 0 {
		return http.StatusNotFound, common.ARTICLE_NOT_FOUND, data
	}
	_ = utils.RedisSet(fmt.Sprintf(QUERY_ARTICLE_BY_ID_KEY, id), data)
	return http.StatusOK, common.SUCCESS, data
}

func GetArticles(pageNum int) (int, int, []ArticleVO) {
	var data []ArticleVO
	//if articles, err := utils.RedisGet(fmt.Sprintf(QUERY_ARTICLES_BY_PAGE_KEY, pageNum - 1)); err == nil{
	//	err = c.Json.Unmarshal(utils.Str2bytes(articles), &data)
	//	if err == nil {
	//		return http.StatusOK, common.SUCCESS, data
	//	}
	//}
	if err := c.Db.
		Raw("SELECT article.id, title, abstract, post_time, author_id, visit_count, last_modify_time, label_name, description "+
			"FROM article INNER JOIN label on article.label_id = label.id "+
			"WHERE article.status > 0 "+
			"ORDER BY post_time DESC LIMIT ?, 5", (pageNum-1)*5).
		Scan(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusNotFound, common.ARTICLE_NOT_FOUND, nil
		}
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get Articles (database)", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	//_ = utils.RedisSet(fmt.Sprintf(QUERY_ARTICLES_BY_PAGE_KEY, pageNum - 1), data)
	return http.StatusOK, common.SUCCESS, data
}

func ModifyArticle(id int64, data *Article) (int, int) {
	if data.Title == "" {
		return http.StatusBadRequest, common.EMPTY_VISITOR_INFO
	}
	var m = make(map[string]interface{})
	m["title"] = data.Title
	m["abstract"] = data.Abstract
	m["content"] = data.Content
	m["label_id"] = data.LabelId
	m["last_modify_time"] = nil
	if err := utils.RedisDelete(fmt.Sprintf(QUERY_ARTICLE_BY_ID_KEY, id)); err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete Article cache (redis)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	if err := c.Db.Model(&data).Where("id = ?", id).Updates(m).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to modify Article (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusOK, common.SUCCESS
}

func DeleteArticle(id int64) (int, int) {
	var m = make(map[string]interface{})
	var article Article
	m["status"] = 0
	if err := utils.RedisDelete(fmt.Sprintf(QUERY_ARTICLE_BY_ID_KEY, id)); err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete Article cache (redis)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	err := c.Db.Model(&article).Where("id = ?", id).Updates(m).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to disable Article (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusOK, common.SUCCESS
}

func GetArticlesByLabel(id int64) (int, int, []Article) {
	var label Label
	if err := c.Db.Where("id = ?", id).First(&label).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get Articles (database); step1", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	if err := c.Db.Model(&label).Related(&label.Articles).
		Select("id, title, abstract, post_time, author_id, visit_count, last_modify_time").
		Where("status > ?", 0).
		Find(&label.Articles).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get Articles (database); step2", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	return http.StatusOK, common.SUCCESS, label.Articles
}

//管理文章时获取文章信息，可以获取到status = 0的文章
func ManageArticles(pageNum int) (int, int, []ArticleVO) {
	var data []ArticleVO
	err := c.Db.
		Raw("SELECT article.id, title, abstract, post_time, author_id, visit_count, last_modify_time, label_name, description, status"+
			"FROM article INNER JOIN label on article.label_id = label.id "+
			"ORDER BY post_time DESC LIMIT ?, 10", (pageNum-1)*10).
		Scan(&data).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusNotFound, common.ARTICLE_NOT_FOUND, nil
		}
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get Articles (database)", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	return http.StatusOK, common.SUCCESS, data
}

func EnableArticle(id int64) (int, int) {
	var m = make(map[string]interface{})
	var article Article
	m["status"] = 1
	err := c.Db.Model(&article).Where("id = ?", id).Updates(m).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to enable Article (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusOK, common.SUCCESS
}
