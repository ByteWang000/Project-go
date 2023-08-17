package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/server/proto"
	"net"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("服务器SayHello收到:", req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil

	//return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func main() {
	//
	listen, _ := net.Listen("tcp", ":9090")
	//创建grpc服务
	grpcServer := grpc.NewServer()

	// 在gprc服务端注册编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务，代码在这里运行，不会向下运行
	grpcerr := grpcServer.Serve(listen)

	if grpcerr != nil {
		fmt.Printf("failed to serve: %v", grpcerr)
		return
	}
}
