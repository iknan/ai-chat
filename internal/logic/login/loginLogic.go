package login

import (
	"ai_chat/internal/common/errNo"
	"ai_chat/internal/infra/mysql/model"
	"ai_chat/internal/infra/redis"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logc"
	"gorm.io/gorm"
	"time"

	"ai_chat/internal/svc"
	"ai_chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	userCache *redis.UserCache
	userModel *model.UserModel
}

// 登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		userCache: redis.NewUserCache(ctx, svcCtx.Redis),
		userModel: model.NewUserModel(svcCtx.DB),
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	var accessToken string

	resp = &types.LoginResp{}
	if req.Phone == "" {
		err = errNo.ReturnRespErr(errNo.ErrParam)
		return
	}
	// 校验验证码和手机号
	errCode := l.verifyCode(req.Phone, req.Code)
	if errCode != 0 {
		err = errNo.ReturnRespErr(errCode)
		return
	}

	// 老用户就直接获取token，新用户自动注册
	userByPhone, errCode := l.isOldUser(req.Phone)
	if errCode != 0 {
		err = errNo.ReturnRespErr(errCode)
		logc.Error(l.ctx, logx.Field("path", "/v1/user/login.Login"), err)
		return
	}

	accessToken, err = l.GenToken(map[string]interface{}{"uid": userByPhone.ID})
	if err != nil {
		err = errNo.ReturnRespErr(errNo.ErrGetToken)
		logc.Error(l.ctx, logx.Field("path", "/v1/user/login.Login"), err)
		return
	}
	l.userCache.SetTokenCache(accessToken)

	resp = &types.LoginResp{
		Token:  accessToken,
		UserId: userByPhone.ID,
	}
	return
}

func (l *LoginLogic) GenToken(payloads map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + l.svcCtx.Config.JwtAuth.AccessExpire
	claims["iat"] = iat
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
	if err != nil {
		return "", err
	}

	return "Bearer " + signedToken, nil
}

// 校验验证码和手机号
func (l *LoginLogic) verifyCode(phone, code string) (errCode int) {
	errCode = 0
	result, err := l.svcCtx.Redis.Get(l.ctx, redis.USER_VERIFY_CODE+phone).Result()
	if err != nil {
		errCode = errNo.ErrGetVerifyCode
		logc.Error(l.ctx, logx.Field("path", "/v1/user/login.Login"), logx.Field("err : ", err), errNo.ErrFactory(errCode))
		return
	}
	if result != code {
		errCode = errNo.ErrVerifyCode
		logc.Error(l.ctx, logx.Field("path", "/v1/user/login.Login"), logx.Field("err : ", err), errNo.ErrFactory(errCode))
		return
	}
	//校验成功后，将验证码删除
	l.svcCtx.Redis.Del(l.ctx, redis.USER_VERIFY_CODE+phone)

	return
}

func (l *LoginLogic) isOldUser(phone string) (userByPhone *model.User, errCode int) {
	var err error
	errCode = 0
	userByPhone, err = l.userModel.GetUserByPhone(phone)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errCode = errNo.ErrDBGetUser
		return
	}

	// 新用户
	if userByPhone.ID == 0 {
		userByPhone = &model.User{
			Phone:    phone,
			Username: phone,
			Avatar:   "http://uplifeting-web-oss.oss-cn-shanghai.aliyuncs.com/web/default_avatar-1740919645851-x13q2.png",
		}
		err = l.userModel.InsertUser(userByPhone)
		if err != nil {
			errCode = errNo.ErrUserRegister
			return
		}
	}

	return
}
