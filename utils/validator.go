package utils

import (
	"MarvelousBlog-Backend/config"
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) (string, bool) {
	err := config.ProjectValidator.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(config.Trans), false
		}
	}
	return "", true
}
