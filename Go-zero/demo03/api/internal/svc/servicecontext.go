package svc

import (
	"demo03/api/internal/config"
	"demo03/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpc)),
	}
}
