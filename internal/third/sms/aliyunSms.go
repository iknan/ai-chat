package sms

import (
	"ai_chat/internal/svc"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/zeromicro/go-zero/core/logx"
)

var phoneRegex = regexp.MustCompile(`^1\d{10}$`)

func init() {
	rand.Seed(time.Now().UnixNano()) // 只执行一次
}

func SendSms(svcCtx *svc.ServiceContext, phone string) (string, error) {
	// 严格校验手机号格式
	if !phoneRegex.MatchString(phone) {
		logx.Errorf("invalid phone number: %s", phone)
		return "", errors.New("手机号格式不正确")
	}

	client, err := dysmsapi.NewClientWithAccessKey(
		svcCtx.Config.AliYun.Region,
		svcCtx.Config.AliYun.AccessKeyId,
		svcCtx.Config.AliYun.AccessKeySecret,
	)
	if err != nil {
		return "", errors.New("短信发送客户端获取失败")
	}

	code := generateVerificationCode()

	// 创建发送短信请求
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = "iknan"             // 替换为你的短信签名
	request.TemplateCode = "SMS_465334254" // 替换为你的短信模板CODE
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code)

	response, err := client.SendSms(request)
	if err != nil {
		logx.Errorf("短信发送失败: %v", err)
		return "", fmt.Errorf("短信发送失败: %w", err)
	}
	if response.Code != "OK" {
		return "", fmt.Errorf("短信发送失败: %s", response.Message)
	}

	return code, nil
}

// 生成 6 位数字验证码，且第一位不为 0
func generateVerificationCode() string {
	firstDigit := rand.Intn(9) + 1   // 确保第一位不为 0
	otherDigits := rand.Intn(100000) // 生成 5 位数字，确保总长度为 6
	return fmt.Sprintf("%d%05d", firstDigit, otherDigits)
}
