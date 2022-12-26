package handler

import (
	"gozerotserver/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozerotserver/app/user/api/internal/logic"
	"gozerotserver/app/user/api/internal/svc"
	"gozerotserver/app/user/api/internal/types"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLogin
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		response.Response(w, resp, err)
	}
}
