package service

import (
	"MarvelousBlog-Backend/common"
	c "MarvelousBlog-Backend/config"
	. "MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/utils"
	"net/http"
)

const ALL_LABELS_KEY = "ALL_LABELS"

//查询标签是否重名
func labelUsed(name string) bool {
	var label Label
	c.Db.Select("id").Where("label_name = ?", name).First(&label)
	//如果查询到的label id > 0， 即证明已经有了这个标签
	if label.Id > 0 {
		//昵称已被占用
		return true
	}
	return false
}

//新增标签
func CreateLabel(data *Label) (int, int) {
	if data.LabelName == "" {
		return http.StatusBadRequest, common.EMPTY_LABEL_INFO
	}
	if labelUsed := labelUsed(data.LabelName); labelUsed {
		return http.StatusForbidden, common.LABEL_USED
	}
	if err := utils.RedisDelete(ALL_LABELS_KEY); err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete old key (redis)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	if err := c.Db.Create(&data).Error; err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to create Label (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusCreated, common.SUCCESS
}

//获取全部标签
func GetAllLabels() (int, int, []Label) {
	var data []Label
	labels, err := utils.RedisGet(ALL_LABELS_KEY)
	if err == nil {
		err = c.Json.Unmarshal(utils.Str2bytes(labels), &data)
		if err == nil {
			return http.StatusOK, common.SUCCESS, data
		}
	}
	err = c.Db.Find(&data).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to get labels (database)", err)
		return http.StatusInternalServerError, common.FAIL, nil
	}
	_ = utils.RedisSet(ALL_LABELS_KEY, data)
	return http.StatusOK, common.SUCCESS, data
}

//修改标签信息
func ModifyLabel(id int64, data *Label) (int, int) {
	var m = make(map[string]interface{})
	m["label_name"] = data.LabelName
	m["description"] = data.Description
	if err := utils.RedisDelete(ALL_LABELS_KEY); err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete old key (redis)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	err := c.Db.Model(&data).Where("id = ?", id).Updates(m).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to modify label (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusOK, common.SUCCESS
}

//删除标签
func DeleteLabel(id int64) (int, int) {
	var label Label
	if err := utils.RedisDelete(ALL_LABELS_KEY); err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete old key (redis)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	err := c.Db.Where("id = ?", id).Delete(&label).Error
	if err != nil {
		c.Log.Errorf(common.SYSTEM_ERROR_LOG, "Fail to delete label (database)", err)
		return http.StatusInternalServerError, common.FAIL
	}
	return http.StatusNoContent, common.SUCCESS
}
