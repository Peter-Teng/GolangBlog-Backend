package config

import (
	"MarvelousBlog-Backend/common"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/sirupsen/logrus"
	"reflect"
)

var ProjectValidator *validator.Validate
var Trans unTrans.Translator

func init() {
	ProjectValidator = validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	Trans, _ = uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(ProjectValidator, Trans)
	if err != nil {
		logrus.Errorf(common.SYSTEM_ERROR_LOG, "validator initialized error", err)
	}
	ProjectValidator.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
}
