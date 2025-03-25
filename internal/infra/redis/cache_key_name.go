package redis

import "time"

const TokenExpireTime = time.Second * 2592000

const (
	KeyUserToken     = "userToken:%s"
	USER_VERIFY_CODE = "user_verify_code:"
)
