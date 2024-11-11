package main

import (
	"computer_network/go-grpc-example/pkg"
	"computer_network/go-grpc-example/proto/token"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"time"
)

type TokenService struct {
	// 必须嵌入UnimplementedUserServiceServer
	token.UnimplementedTokenServiceServer
}

// 实现SayHi方法
func (t *TokenService) Token(ctx context.Context, req *token.Request) (res *token.Response, err error) {
	format := time.Now().Format("2006-01-02 15:04:05")
	return &token.Response{Name: "hi " + req.GetName() + "---" + format, Uid: 1}, nil
}

const PORT = "8888"

func main() {
	// 监听端口
	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	var opts []grpc.ServerOption

	if pkg.IsTLS {
		// TLS认证
		// 根据服务端输入的证书文件和密钥构造 TLS 凭证
		c, err := credentials.NewServerTLSFromFile("./go-grpc-example/conf/server.pem", "./go-grpc-example/conf/server.key")
		if err != nil {
			log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
		}
		opts = append(opts, grpc.Creds(c))
	}

	opts = append(opts, grpc.UnaryInterceptor(pkg.ServerInterceptorCheckToken()))
	// 创建grpc服务
	server := grpc.NewServer(opts...)
	// 注册服务
	token.RegisterTokenServiceServer(server, &TokenService{})

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}
