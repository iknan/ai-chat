package login

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"ai_chat/internal/logic/login"
	"ai_chat/internal/svc"
	"ai_chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 登录
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)

		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}

	}
}
