// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	article "gozero/app/article/cmd/api/internal/handler/article"
	"gozero/app/article/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/getArticles",
				Handler: article.GetArticlesHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/createArticle",
				Handler: article.CreateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/deleteArticle",
				Handler: article.DeleteArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateArticle",
				Handler: article.UpdateArticleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/article/v1"),
	)
}