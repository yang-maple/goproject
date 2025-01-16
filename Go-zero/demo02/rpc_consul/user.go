package main

import (
	"flag"
	"fmt"

	"demo02/rpc_consul/internal/config"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		_ = consul.RegisterService(c.ListenOn, c.Consul)
		// pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		// if c.Mode == service.DevMode || c.Mode == service.TestMode {
		// 	reflection.Register(grpcServer)
		// }
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
