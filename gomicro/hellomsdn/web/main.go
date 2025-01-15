package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	pb "hellomsdn/proto"
)

func main() {

	service := micro.NewService()
	service.Init()
	//pb.NewHellomsdnService("hellomsdn", service.Client()) 该函数需要到 hellomsdn.pb.micro.go 中找
	client := pb.NewHellomsdnService("hellomsdn", service.Client())
	rsp, err := client.Call(context.TODO(), &pb.Request{
		Name: "hello go micro in hellomsdn",
	})
	if err != nil {
		panic(err)
	}
	println(rsp.Msg)
}
