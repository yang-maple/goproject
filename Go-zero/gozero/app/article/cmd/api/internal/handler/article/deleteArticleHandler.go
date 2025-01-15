package article

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gozero/app/article/cmd/api/internal/logic/article"
	"gozero/app/article/cmd/api/internal/svc"
	"gozero/app/article/cmd/api/internal/types"
)

func DeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := article.NewDeleteArticleLogic(r.Context(), svcCtx)
		resp, err := l.DeleteArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
