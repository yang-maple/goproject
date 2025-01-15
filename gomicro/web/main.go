package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/gin-gonic/gin"

	pb "web/proto"
)

func main() {

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		service := micro.NewService()
		service.Init()
		//pb.NewHellomsdnService("hellomsdn", service.Client()) 该函数需要到 hellomsdn.pb.micro.go 中找
		client := pb.NewHellomsdnService("hellomsdn", service.Client())
		rsp, err := client.Call(context.TODO(), &pb.Request{
			Name: "hello go micro in 123445",
		})
		if err != nil {
			panic(err)
		}
		println(rsp.Msg)
		c.JSON(200, gin.H{
			"message": rsp.Msg,
		})
	})
	router.Run(":8000")
}
