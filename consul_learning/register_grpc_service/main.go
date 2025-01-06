package main

import (
	"computer_network/consul_learning/consul"
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
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
	server := grpc.NewServer()
	// 注册服务
	hello.RegisterUserServiceServer(server, &HelloService{})
	//将grpc服务注册到consul
	cc, err := consul.NewConsul("127.0.0.1:8500")
	if err != nil {
		log.Fatalf("consul.NewConsul err: %v", err)
		return
	}

	//gRPC服务支持健康检查,记得导入对应包
	healthCheck := health.NewServer()
	healthpb.RegisterHealthServer(server, healthCheck)

	serverIP, err := consul.GetOutboundIP()
	if err != nil {
		log.Fatalf("GetOutboundIP err: %v", err)
		return
	}

	if err := cc.RegisterService("grpc-learing2", serverIP.String(), 8888); err != nil {
		log.Fatalf("consul.RegisterService err: %v", err)
		return
	}

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
