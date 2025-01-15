package main

import (
	"context"
	"demo4/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 使用安全连接 连接grpc 服务端
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("grpc 连接失败" + err.Error())
	}
	defer conn.Close()
	// 创建gRPC客户端
	client := pb.NewComputeClient(conn)

	// 调用gRPC服务端 Area方法
	Areq, err := client.Area(context.Background(), &pb.AreaRequest{Length: 5, Width: 10})
	if err != nil {
		panic(err)
	}
	fmt.Println(Areq.String())
	// 调用gRPC服务端 Perimeter方法
	Preq, err := client.Perimeter(context.Background(), &pb.PerimeterRequest{Length: 5, Width: 10})
	if err != nil {
		panic(err)
	}
	println(Preq.String())
}
