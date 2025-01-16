package main

import (
	"flag"
	"fmt"

	"demo01/internal/config"
	"demo01/internal/server"
	"demo01/internal/svc"
	"demo01/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/article.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterArticleServer(grpcServer, server.NewArticleServer(ctx))

		// if c.Mode == service.DevMode || c.Mode == service.TestMode {
		// 	reflection.Register(grpcServer)
		// }
	})
	_ = consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
