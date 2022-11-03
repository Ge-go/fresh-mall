package utils

import (
	"fmt"
	"mall-api/user-web/global"
	"math/rand"
	"time"
	"unsafe"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

const letterBytes = "0123456789"

var src = rand.NewSource(time.Now().UnixNano())

// GenerateSmsCode 生成n位验证码
func GenerateSmsCode(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// SendSms 阿里云提供短信服务,仅用于自己账号测试
func SendSms(code string) error {
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential(global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecret)
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := dysmsapi.NewClientWithOptions(global.ServerConfig.AliSmsInfo.RegionId, config, credential)
	if err != nil {
		panic(err)
	}

	request := dysmsapi.CreateSendSmsRequest()

	request.Scheme = "https"

	request.SignName = "阿里云短信测试"
	request.TemplateCode = "SMS_154950909"
	request.PhoneNumbers = "15093710052"
	template := fmt.Sprintf("{\"code\":%q}", code)
	request.TemplateParam = template

	_, err = client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}
