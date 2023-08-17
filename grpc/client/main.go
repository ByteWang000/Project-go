package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/server/proto"
	"log"
)

func main() {

	// 连接到服务端，此处禁用安全传输，无加密和验证
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}

	defer conn.Close()

	client := pb.NewSayHelloClient(conn)
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "wangzijie"})
	//fmt.Println("这是客户端收到消息：", resp)
	fmt.Println("这是客户端解析：", resp.ResponseMsg)
}
