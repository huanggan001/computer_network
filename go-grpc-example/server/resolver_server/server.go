package main

import (
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type HelloService struct {
	// 必须嵌入UnimplementedUserServiceServer
	hello.UnimplementedUserServiceServer
	addr string
}

var port = flag.String("port", "8972", "The server port")

// 实现SayHi方法
func (h *HelloService) Say(ctx context.Context, req *hello.Request) (res *hello.Response, err error) {
	format := time.Now().Format("2006-01-02 15:04:05")
	return &hello.Response{Result: "hi " + req.GetName() + "---" + format + "from " + h.addr}, nil
}

func main() {
	flag.Parse()
	// 监听端口
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	hello.RegisterUserServiceServer(server, &HelloService{addr: *port})

	fmt.Println("server listening at ", *port)
	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
