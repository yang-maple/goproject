package main

import (
	"context"
	pb "github.com/go-micro/examples/helloworld/proto"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/hashicorp/consul/api"
	micro "go-micro.dev/v4"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	consulReg := consul.NewRegistry(
		consul.Config(&api.Config{
			Address: "10.1.131.31:8500",
			Token:   "74564183-951d-27e3-493e-bb57bfbb3892",
		}),
	)
	service := micro.NewService(
		micro.Registry(consulReg),
		micro.Name("helloworld"),
		micro.Address(":8080"),
	)

	service.Init()

	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
