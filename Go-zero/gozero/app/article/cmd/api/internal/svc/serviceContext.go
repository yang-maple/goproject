package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gozero/app/article/cmd/api/internal/config"
	"gozero/app/article/cmd/rpc/article"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRpc article.ArticleZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ArticleRpc: article.NewArticleZrpcClient(zrpc.MustNewClient(c.ArticleRpcConf)),
	}
}
