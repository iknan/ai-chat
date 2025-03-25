package errNo

import (
	"fmt"
	"github.com/zeromicro/x/errors"
)

const (
	Success = 0

	ErrParam         = 10001
	ErrGetToken      = 10002
	ErrSmsSendFailed = 10008
	ErrGetVerifyCode = 10009
	ErrVerifyCode    = 10010
	ErrTokenInvalid  = 10011
	ErrDBGetUser     = 10012
	ErrUserRegister  = 10013
)

var errMap = map[int]string{
	Success:          "成功",
	ErrParam:         "参数错误",
	ErrGetToken:      "获取token失败",
	ErrSmsSendFailed: "短信发送失败",
	ErrGetVerifyCode: "获取验证码失败",
	ErrVerifyCode:    "验证码错误",
	ErrTokenInvalid:  "token无效",
	ErrDBGetUser:     "数据库获取用户失败",
	ErrUserRegister:  "用户注册失败",
}

func ErrFactory(errCode int) error {
	return fmt.Errorf("code : %d , msg : %s", errCode, errMap[errCode])
}

func ReturnRespErr(errCode int) error {
	return errors.New(errCode, errMap[errCode])
}
