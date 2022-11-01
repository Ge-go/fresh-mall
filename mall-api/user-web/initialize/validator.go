package initialize

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"

	"mall-api/user-web/global"
	mallValidator "mall-api/user-web/validator"
)

// InitTrans 初始化 validator 验证器
func InitTrans(locale string) (err error) {
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			_ = entranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zhtranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = entranslations.RegisterDefaultTranslations(v, global.Trans)
		}

		// 自定义Validator mobile
		if vld, ok := binding.Validator.Engine().(*validator.Validate); ok {
			_ = vld.RegisterValidation("mobile", mallValidator.ValidateMobile)
			_ = vld.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
				return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("mobile", fe.Field())
				return t
			})
		}

		return
	}

	return
}
