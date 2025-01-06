package main

import (
	"context"
	"crypto/tls"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/bookstore/data"
	books "grpc/bookstore/pb"
	"grpc/bookstore/sql"
	"grpc/pkg/util"
	"log"
	"net"
	"net/http"
)

// 使用一个端口提供gprc和https服务
func main() {
	//创建tcp连接
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sql.InitMySql()
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	var ops []grpc.ServerOption
	creds, err := credentials.NewClientTLSFromFile("./cert/server.pem", "./cert/server.key")
	if err != nil {
		log.Fatalf("failed to init cert: %v", err)
	}
	ops = append(ops, grpc.Creds(creds))

	//创建了一个没有注册服务的grpc服务端
	grpcServer := grpc.NewServer(ops...)

	//注册grpc服务
	books.RegisterBookstoreServer(grpcServer, &data.BookStore{DB: db})
	// 创建 grpc-gateway 关联组件
	// context.Background()返回一个非空的空上下文。
	// 它没有被注销，没有值，没有过期时间。它通常由主函数、初始化和测试使用，并作为传入请求的顶级上下文
	ctx := context.Background()
	//从客户端的输入证书文件构造TLS凭证
	dcreds, err := credentials.NewClientTLSFromFile("./cert/server.pem", "grpc_example")
	if err != nil {
		log.Printf("Failed to create client TLS credentials %v", err)
	}
	// grpc.WithTransportCredentials 配置一个连接级别的安全凭据(例：TLS、SSL)，返回值为type DialOption
	// grpc.DialOption DialOption选项配置我们如何设置连接（其内部具体由多个的DialOption组成，决定其设置连接的内容）
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	// 创建HTTP NewServeMux及注册grpc-gateway逻辑
	// runtime.NewServeMux：返回一个新的ServeMux，它的内部映射是空的；
	// ServeMux是grpc-gateway的一个请求多路复用器。它将http请求与模式匹配，并调用相应的处理程序
	gwmux := runtime.NewServeMux()
	// RegisterHelloWorldHandlerFromEndpoint：注册HelloWorld服务的HTTP Handle到grpc端点
	if err := books.RegisterBookstoreHandlerFromEndpoint(ctx, gwmux, "127.0.0.1:8080", dopts); err != nil {
		log.Printf("Failed to register_grpc_service gw server: %v\n", err)
	}
	//http服务
	//分配并返回一个新的ServeMux
	mux := http.NewServeMux()
	//为给定模式注册处理模式
	mux.Handle("/", gwmux)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: util.GrpcHandlerFunc(grpcServer, mux),
	}
	// 通过util.GetTLSConfig解析得到tls.Config，传达给http.Server服务的TLSConfig配置项使用
	tlsConfig := util.GetTLSConfig("./cert/server.pem", "./cert/server.key")
	log.Printf("grpc and https listen on : 8080")
	log.Fatal(srv.Serve(tls.NewListener(lis, tlsConfig)))
}
