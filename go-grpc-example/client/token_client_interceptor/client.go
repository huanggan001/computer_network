package main

import (
	"computer_network/go-grpc-example/pkg"
	"computer_network/go-grpc-example/proto/token"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = "8888"

func main() {
	// 建立链接
	//conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(pkg.UnaryClientInterceptor()))
	var opts []grpc.DialOption

	auth := pkg.TokenAuth{
		token.TokenValidateParam{
			Token: "81dc9bdb52d04dc20036dbd8313ed055",
			Uid:   1234,
		},
	}
	if auth.RequireTransportSecurity() {
		//打开tls 走tls认证
		// 根据客户端输入的证书文件和密钥构造 TLS 凭证。
		// 第二个参数 serverNameOverride 为服务名称。
		c, err := credentials.NewClientTLSFromFile("./go-grpc-example/conf/server.pem", "go-grpc-example")
		if err != nil {
			log.Fatalf("credentials.NewClientTLSFromFile err: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(c))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	opts = append(opts, grpc.WithPerRPCCredentials(&auth))

	//链式拦截器
	conn, err := grpc.Dial("localhost:8888", opts...)
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	// 一定要记得关闭链接
	defer conn.Close()
	// 实例化客户端
	client := token.NewTokenServiceClient(conn)
	// 发起请求
	response, err := client.Token(context.Background(), &token.Request{Name: "ss"})
	if err != nil {
		log.Fatalf("client.Token err: %v", err)
		//fmt.Printf("Say err: %v", err)
	}
	fmt.Println("============")
	fmt.Printf("resp: %s \n", response.String())
	fmt.Println("============")

}
