package logic

import (
	"context"

	"ai_chat/internal/svc"
	"ai_chat/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Ai_chatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAi_chatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Ai_chatLogic {
	return &Ai_chatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Ai_chatLogic) Ai_chat(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
