package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/bookstore/data"
	books "grpc/bookstore/pb"
	"grpc/bookstore/sql"
	"log"
	"net"
	"net/http"
)

func main() {
	//创建tcp连接
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创一个grpc server
	grpcServer := grpc.NewServer()
	db, err := sql.InitMySql()
	if err != nil {
		log.Fatalf("failed to init db: %v", err)
	}
	//注册bookstore service到server
	books.RegisterBookstoreServer(grpcServer, &data.BookStore{DB: db})
	// 8080端口启动gRPC Server
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		log.Fatal(grpcServer.Serve(lis))
	}()
	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	gwmux := runtime.NewServeMux()
	//注册bookstore
	err = books.RegisterBookstoreHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	gwServer := &http.Server{
		Addr:    ":8081",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8081")
	log.Fatalln(gwServer.ListenAndServe())
}
