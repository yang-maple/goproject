package handler

import (
	"net/http"

	"demo03/api/internal/logic"
	"demo03/api/internal/svc"
	"demo03/api/internal/types"
	"demo03/common/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func deleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDeleteLogic(r.Context(), svcCtx)
		resp, err := l.Delete(&req)
		response.Response(r, w, resp, err)
	}
}
