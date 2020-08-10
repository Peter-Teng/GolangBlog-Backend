package service

import (
	"MarvelousBlog-Backend/common"
	c "MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity/model"
	"github.com/jinzhu/gorm"
	"net/http"
)

func MakeComment(data *model.Comment) (int, int) {
	if err := c.Db.Create(&data).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to create Article (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusCreated, common.SUCCESS
}

func ListComment(id int64, pageNum int) (int, int, []model.Comment) {
	var data []model.Comment
	if err := c.Db.Limit(20).Offset((pageNum-1)*20).Where("article_id = ?", id).Find(&data).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return http.StatusNotFound, common.USER_NOT_FOUND, nil
		}
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to load comments (database)", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	return http.StatusOK, common.SUCCESS, data
}

func DeleteComment(id int64) (int, int) {
	var comment model.Comment
	if err := c.Db.Where("id = ?", id).Delete(&comment).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete visitor (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusNoContent, common.SUCCESS
}
