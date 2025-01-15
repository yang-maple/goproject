package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	pb "helloweb/proto"
)

func main() {
	r := gin.Default()
	r.GET("/hello", CallRemoter)
	r.Run(":8010")
	// Create service

}

func CallRemoter(c *gin.Context) {
	consulReg := consul.NewRegistry(
		consul.Config(&api.Config{
			Address: "10.1.131.31:8500",
			Token:   "74564183-951d-27e3-493e-bb57bfbb3892",
		}),
	)
	service := micro.NewService(
		micro.Registry(consulReg),
	)
	service.Init()
	client := pb.NewHellowordService("helloword", service.Client())
	// Register handler
	rsp, err := client.Call(context.TODO(), &pb.Request{
		Name: "micro445566777",
	})

	if err != nil {
		panic(err)
	}
	println(rsp.Msg)
	c.JSON(200, gin.H{
		"message": rsp.Msg,
	})
}
