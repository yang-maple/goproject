package handler

import (
	"net/http"

	"demo02/api_jwt/internal/logic"
	"demo02/api_jwt/internal/svc"
	"demo02/common/response"

	_ "github.com/zeromicro/go-zero/rest/httpx"
)

func userInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo()
		response.Response(r, w, resp, err)
	}
}
