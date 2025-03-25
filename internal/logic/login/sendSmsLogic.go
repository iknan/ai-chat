package login

import (
	"ai_chat/internal/common/errNo"
	"ai_chat/internal/infra/redis"
	"ai_chat/internal/third/sms"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"ai_chat/internal/svc"
	"ai_chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	userCache *redis.UserCache
}

// 发送验证码
func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		userCache: redis.NewUserCache(ctx, svcCtx.Redis),
	}
}

func (l *SendSmsLogic) SendSms(req *types.SendSmsReq) (err error) {
	DurationCmd := l.svcCtx.Redis.TTL(l.ctx, redis.USER_VERIFY_CODE+req.Phone)
	ttl := DurationCmd.Val()
	if ttl > 3 {
		err = errNo.ReturnRespErr(errNo.ErrGetVerifyCode)
		logc.Error(l.ctx, logx.Field("path", "/user/login.SendSms"), errNo.ErrFactory(errNo.ErrGetVerifyCode))
		return
	}
	code, err := sms.SendSms(l.svcCtx, req.Phone)

	if err != nil {
		err = errNo.ReturnRespErr(errNo.ErrSmsSendFailed)
		logc.Error(l.ctx, logx.Field("path", "/user/login.SendSms"), errNo.ErrFactory(errNo.ErrSmsSendFailed))
		return
	}
	// 发送成功要存到redis，五分钟过期
	l.svcCtx.Redis.Set(l.ctx, redis.USER_VERIFY_CODE+req.Phone, code, 5*time.Minute)

	return
}
