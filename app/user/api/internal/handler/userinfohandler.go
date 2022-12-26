package handler

import (
	"gozerotserver/common/response"
	"net/http"

	"gozerotserver/app/user/api/internal/logic"
	"gozerotserver/app/user/api/internal/svc"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(w, resp, err)
	}
}
