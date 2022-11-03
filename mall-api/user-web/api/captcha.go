package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	_const "mall-api/user-web/const"
	"mall-api/user-web/forms"
	"mall-api/user-web/global"
	"mall-api/user-web/utils"
	"net/http"
	"time"
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
	//表单验证
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&sendSmsForm); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	// 生成6位验证码
	smsCode := utils.GenerateSmsCode(6)

	// send
	if err := utils.SendSms(smsCode); err != nil {
		zap.S().Errorw("send sms error", "msg", err.Error())
		return
	}

	var key string
	// register or login
	if sendSmsForm.Type == 1 { // register
		// business prefix
		key = fmt.Sprintf("%s%s", _const.SmsRegisterPrefix, sendSmsForm.Mobile)
	} else { // login
		key = fmt.Sprintf("%s%s", _const.SmsLoginPrefix, sendSmsForm.Mobile)
	}

	//set redis
	cmd := global.RedisClient.Set(ctx, key, smsCode, 5*time.Minute)
	if cmd.Err() != nil {
		zap.S().Errorw("set redis", "msg", cmd.Err())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "send success",
	})
}
