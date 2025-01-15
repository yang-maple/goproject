package main

import (
	"context"
	"demo4/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type computeServer struct{}

// Area 实现计算面积的方法
func (*computeServer) Area(ctx context.Context, req *pb.AreaRequest) (*pb.AreaResponse, error) {
	return &pb.AreaResponse{
		Area: req.Length * req.Width,
	}, nil
}

// Perimeter 实现计算周长的方法
func (*computeServer) Perimeter(ctx context.Context, req *pb.PerimeterRequest) (*pb.PerimeterResponse, error) {
	return &pb.PerimeterResponse{
		Perimeter: 2 * (req.Length + req.Width),
	}, nil
}

func main() {
	//创建gRPC服务器实例
	grpcServer := grpc.NewServer()
	//注册gRPC服务
	pb.RegisterComputeServer(grpcServer, new(computeServer))
	//监听端口
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server is running...")
	//启动gRPC服务器
	grpcServer.Serve(lis)
}
