package main

import (
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"sync"
	"time"
)

type HelloService struct {
	// 必须嵌入UnimplementedUserServiceServer
	vis map[string]struct{}
	//添加并发锁
	mu sync.Mutex
	hello.UnimplementedUserServiceServer
}

// 实现SayHi方法
// todo 添加限制，每个name只能请求一次gprc服务
func (h *HelloService) Say(ctx context.Context, req *hello.Request) (res *hello.Response, err error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if _, ok := h.vis[req.GetName()]; ok {
		err := status.Errorf(codes.AlreadyExists, "每个name只能访问一次")
		return nil, err
	}
	h.vis[req.GetName()] = struct{}{}
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
	server := grpc.NewServer()
	// 注册服务
	hello.RegisterUserServiceServer(server, &HelloService{vis: make(map[string]struct{})})

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
