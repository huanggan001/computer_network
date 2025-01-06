package main

import (
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const PORT = "8888"

func main() {
	// 建立链接
	//conn, err := grpc.Dial("dns:///localhost:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 使用自定义的resolver
	//conn, err := grpc.Dial(
	//	"q1mi:///resolver.liwenzhou.com",
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	conn, err := grpc.Dial(
		"q1mi:///resolver.liwenzhou.com",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithResolvers(&q1miResolverBuilder{}), // 指定使用q1miResolverBuilder
	)

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	// 一定要记得关闭链接
	defer conn.Close()
	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	//response, err := client.Say(context.Background(), &hello.Request{Name: "ss"})
	//if err != nil {
	//	log.Fatalf("client.Say err: %v", err)
	//	//fmt.Printf("Say err: %v", err)
	//}
	//fmt.Printf("resp: %s", response.String())
	// 调用RPC方法
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		resp, err := client.Say(ctx, &hello.Request{Name: "ss"})
		if err != nil {
			fmt.Printf("c.SayHello failed, err:%v\n", err)
			return
		}
		// 拿到了RPC响应
		fmt.Printf("resp:%v\n", resp.String())
	}
}
