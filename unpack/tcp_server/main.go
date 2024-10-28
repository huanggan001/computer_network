package main

import (
	"computer_network/unpack/unpack"
	"fmt"
	"net"
)

func main() {

	//监听端口
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go process(conn)
	}

}

func process(conn net.Conn) {
	//defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Println(err)
			break
		}
		str := string(bt)
		fmt.Println("receive from client:", str)
	}
}
