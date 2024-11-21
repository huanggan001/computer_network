package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	//建立连接
	//rpc.Dial 拨号 RPC 服务，然后通过 dial.Call 调用具体的 RPC 方法
	dial, err := rpc.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal("Dial error ", err)
	}
	var result string
	for i := 0; i < 5; i++ {
		//发起请求
		//第一个参数是用点号连接的 RPC 服务名字和方法名字，
		//第二和第三个参数分别我们定义 RPC 方法的两个参数，第一个是客服端传递的消息，第二个是由服务端产生返回的结果。
		_ = dial.Call("HiLinzy.SayHi", "linzy", &result)
		//异步调用为dial.Go
		fmt.Println("rpc service result:", result)
		time.Sleep(time.Second)
	}
}
