package handler

import (
	"net/http"

	"ai_chat/internal/logic"
	"ai_chat/internal/svc"
	"ai_chat/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Ai_chatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAi_chatLogic(r.Context(), svcCtx)
		resp, err := l.Ai_chat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
