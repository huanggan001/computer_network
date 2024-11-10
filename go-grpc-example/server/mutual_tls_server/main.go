package main

import (
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"time"
)

const PORT = "8888"

type HelloService struct {
	// 必须嵌入UnimplementedUserServiceServer
	hello.UnimplementedUserServiceServer
}

// 实现SayHi方法
func (h *HelloService) Say(ctx context.Context, req *hello.Request) (res *hello.Response, err error) {
	format := time.Now().Format("2006-01-02 15:04:05")
	return &hello.Response{Result: "hi " + req.GetName() + "---" + format}, nil
}

func main() {
	// 公钥中读取和解析公钥/私钥对
	cert, err := tls.LoadX509KeyPair("./go-grpc-example/conf/mutual_tls/server.crt", "./go-grpc-example/conf/mutual_tls/server.key")
	if err != nil {
		fmt.Println("LoadX509KeyPair error", err)
		return
	}
	// 创建一组根证书
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("./go-grpc-example/conf/ca.crt")
	if err != nil {
		fmt.Println("read ca pem error ", err)
		return
	}
	// 解析证书
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		fmt.Println("AppendCertsFromPEM error ")
		return
	}

	c := credentials.NewTLS(&tls.Config{
		//设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		//要求必须校验客户端的证书
		ClientAuth: tls.RequireAndVerifyClientCert,
		//设置根证书的集合，校验方式使用ClientAuth设定的模式
		ClientCAs: certPool,
	})
	s := grpc.NewServer(grpc.Creds(c))
	lis, err := net.Listen("tcp", ":"+PORT) //创建 Listen，监听 TCP 端口
	if err != nil {
		log.Fatalf("credentials.NewServerTLSFromFile err: %v", err)
	}
	//将 UserServiceServer（其包含需要被调用的服务端接口）注册到 gRPC Server 的内部注册中心。
	//这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
	hello.RegisterUserServiceServer(s, &HelloService{})

	//gRPC Server 开始 lis.Accept，直到 Stop 或 GracefulStop
	s.Serve(lis)
}
