package main

import (
	"computer_network/go-grpc-gateway/proto"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Reply: in.GetName() + "world"}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//创建一个grpc server对象
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	//启动grpc server
	log.Printf("server listening at %v", listener.Addr())
	go func() {
		log.Fatal(s.Serve(listener))
	}()

	//创建一个连接 连接到我们刚刚启动的grpc服务器的客户端连接
	//grpc-gateway就是通过它来代理请求（将http请求转为grpc请求）
	//阻塞连接，保持同步
	conn, err := grpc.DialContext(context.Background(), "localhost:8080", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	gwmux := runtime.NewServeMux()
	//注册Greeter
	err = proto.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	//开启一个8081的http服务器，充当代理服务器，将http请求转为grpc请求
	gwServer := &http.Server{
		Addr:    ":8081",
		Handler: gwmux,
	}
	// 8081端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8081")
	log.Fatalln(gwServer.ListenAndServe())
}
