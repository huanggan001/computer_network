package main

import (
	"computer_network/unpack/unpack"
	"fmt"
	"net"
	"strconv"
)

func main() {
	//客户端在 TCP 通信中不需要显式指定端口号，而是由操作系统自动分配一个本地端口号来进行通信。
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		fmt.Println("hello tcp_server" + strconv.Itoa(i))
		unpack.Encode(conn, "hello tcp_server"+strconv.Itoa(i))
	}
}
