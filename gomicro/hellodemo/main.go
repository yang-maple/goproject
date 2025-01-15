package main

import (
	"github.com/micro/micro/v3/service"
	"hellodemo/handler"
	pb "hellodemo/proto"

	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	//srv := micro.NewService(
	//	micro.Name("hellodemo"),
	//	micro.Version("latest"),
	//)
	srv := service.New(
		service.Name("hellodemo"),
		service.Address("localhost:8081"),
	)

	// Register handler
	err := pb.RegisterHellodemoHandler(srv.Server(), handler.New())
	if err != nil {
		return
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
