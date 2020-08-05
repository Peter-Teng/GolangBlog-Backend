package utils

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
)

func ParameterWarnLog(err error) {
	config.Log.Warnf(common.SYSTEM_ERROR_LOG, "输入参数错误", err)
}
