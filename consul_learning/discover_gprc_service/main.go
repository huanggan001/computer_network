package main

import (
	"computer_network/go-grpc-example/proto/hello"
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = "8888"

var addr string

// 第一种,通过api的形式来获取grpc服务, 但实际的grpc服务发现不需要这样写
//
//	func main() {
//		//获取consul client
//		cc, err := api.NewClient(api.DefaultConfig())
//		if err != nil {
//			log.Fatalf("api.NewClient failed, err: %v\n", err)
//			return
//		}
//		//根据过滤器的形式获取grpc_learning的服务
//		//有多种获取方式，具体可以查看接口
//		//这里返回的是切片map，因为1个服务可能有多个实例
//		serviceMap, err := cc.Agent().ServicesWithFilter("Service == `grpc-learing2`")
//		if err != nil {
//			log.Fatalf("Agent.ServicesWithFilter failed, err: %v\n", err)
//			return
//		}
//		//选择其中1个实列
//		for k, v := range serviceMap {
//			fmt.Println(k, v)
//			addr = fmt.Sprintf("%s:%s", v.Address, strconv.Itoa(v.Port))
//		}
//
//		// 建立链接
//		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("grpc.Dial err: %v", err)
//		}
//		// 一定要记得关闭链接
//		defer conn.Close()
//		// 实例化客户端
//		client := hello.NewUserServiceClient(conn)
//		// 发起请求
//		response, err := client.Say(context.Background(), &hello.Request{Name: "ss"})
//		if err != nil {
//			log.Fatalf("client.Say err: %v", err)
//			//fmt.Printf("Say err: %v", err)
//		}
//		fmt.Printf("resp: %s", response.String())
//
// }

// 第2种，通过consul resolver来获取grpc服务
func main() {
	// 建立链接
	conn, err := grpc.Dial("consul://127.0.0.1:8500/grpc-learing2?healtht=true", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	// 一定要记得关闭链接
	defer conn.Close()
	// 实例化客户端
	client := hello.NewUserServiceClient(conn)
	// 发起请求
	response, err := client.Say(context.Background(), &hello.Request{Name: "ss"})
	if err != nil {
		log.Fatalf("client.Say err: %v", err)
		//fmt.Printf("Say err: %v", err)
	}
	fmt.Printf("resp: %s", response.String())

}
