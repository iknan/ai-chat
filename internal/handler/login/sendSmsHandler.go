package login

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"ai_chat/internal/logic/login"
	"ai_chat/internal/svc"
	"ai_chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 发送验证码
func SendSmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSmsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewSendSmsLogic(r.Context(), svcCtx)
		err := l.SendSms(&req)
		xhttp.JsonBaseResponseCtx(r.Context(), w, err)
	}
}
