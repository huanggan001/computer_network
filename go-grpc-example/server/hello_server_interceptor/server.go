package main

import (
	"computer_network/go-grpc-example/pkg"
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type HelloService struct {
	// 必须嵌入UnimplementedUserServiceServer
	hello.UnimplementedUserServiceServer
}

// 实现SayHi方法
func (h *HelloService) Say(ctx context.Context, req *hello.Request) (res *hello.Response, err error) {
	format := time.Now().Format("2006-01-02 15:04:05")
	return &hello.Response{Result: "hi " + req.GetName() + "---" + format}, nil
}

const PORT = "8888"

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 创建grpc服务
	server := grpc.NewServer(grpc.UnaryInterceptor(pkg.UnaryServerInterceptor()))
	// 注册服务
	hello.RegisterUserServiceServer(server, &HelloService{})

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
