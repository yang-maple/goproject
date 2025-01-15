package main

import (
	consul "github.com/go-micro/plugins/v4/registry/consul"
	"github.com/hashicorp/consul/api"
	service "go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"helloword/handler"
	pb "helloword/proto"
)

const (
	Address = "10.1.131.31:8500"
	Name    = "helloword"
	Token   = "74564183-951d-27e3-493e-bb57bfbb3892"
)

func main() {
	consulRegistry := consul.NewRegistry(
		consul.Config(&api.Config{
			Address: Address,
			Token:   Token,
		}),
	)

	// Create service

	//srv := service.NewService(
	//	service.Name("helloword"),
	//	service.Version("latest"),
	//)
	srv := service.NewService(
		service.Registry(consulRegistry),
		service.Name(Name),
	)

	// Register handler
	_ = pb.RegisterHellowordHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
