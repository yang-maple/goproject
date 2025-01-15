package main

import (
	"hellomsdn/handler"
	pb "hellomsdn/proto"

	service "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
)

func main() {
	// Create service
	srv := service.NewService(
		service.Name("hellomsdn"),
		service.Version("latest"),
	)

	// Register handler
	err := pb.RegisterHellomsdnHandler(srv.Server(), handler.New())
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
