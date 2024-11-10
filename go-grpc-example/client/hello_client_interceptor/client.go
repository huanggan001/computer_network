package main

import (
	"computer_network/go-grpc-example/pkg"
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = "8888"

func main() {
	// 建立链接
	//conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(pkg.UnaryClientInterceptor()))
	//链式拦截器
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(pkg.UnaryClientInterceptor(), pkg.UnaryClientInterceptorTwo()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	// 一定要记得关闭链接
	defer conn.Close()
	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	response, err := client.Say(context.Background(), &hello.Request{Name: "ss"})
	if err != nil {
		log.Fatalf("client.Say err: %v", err)
		//fmt.Printf("Say err: %v", err)
	}
	fmt.Println("============")
	fmt.Printf("resp: %s \n", response.String())
	fmt.Println("============")

}
