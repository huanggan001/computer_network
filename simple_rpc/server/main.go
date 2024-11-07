package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type HelloService struct{}

// SayHi
// Go 语言的 RPC 规则：方法只能有两个可序列化的参数，
// 其中第二个参数是指针类型，并且返回一个 error 类型，同时必须是公开的方法。
func (h *HelloService) SayHi(request string, response *string) error {
	format := time.Now().Format("2006-01-02 15:04:05")
	*response = "hi " + request + "---" + format
	return nil
}

func main() {
	//注册服务
	//RegisterName类似于Register，但使用提供的名称作为类型，
	//Register 函数调用会将对象类型中所有满足 RPC 规则的对象方法注册为 RPC 函数，所有注册的方法会放在 “HelloService” 服务空间之下。
	_ = rpc.RegisterName("HiLinzy", new(HelloService))
	//监听接口
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
		return
	}
	for {
		//监听请求
		accept, err := lis.Accept()
		if err != nil {
			log.Fatalf("Accept Error: %s", err)
		}
		//用goroutine为每个TCP连接提供RPC服务
		//rpc.ServeConn 函数在该 TCP 连接上为对方提供 RPC 服务。
		//我们的服务支持多个 TCP 连接，然后为每个 TCP 连接提供 RPC 服务。
		go rpc.ServeConn(accept)
	}
}
