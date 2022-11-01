package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateMobile 验证器过滤 正则验证手机号码是否正确
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	//正则过滤Mobile
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)

	if !ok {
		return false
	}
	return true
}
