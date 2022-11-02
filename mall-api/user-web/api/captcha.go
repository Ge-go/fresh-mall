package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"net/http"
)

// 10min有效
var store = base64Captcha.DefaultMemStore

func GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, store)

	id, base64, err := captcha.Generate()
	if err != nil {
		zap.S().Errorw("error generating verification code", "msg", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error generating verification code",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   base64,
	})
}

// SendSms 短信验证
func SendSms(ctx *gin.Context) {

}
