package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var data [1024]byte
		//udp 客户端通信中不需要显式指定端口号，而是由操作系统自动分配一个本地端口号来进行通信。
		n, addr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			fmt.Printf("read failed from addr: %v, err: %v\n", addr, err)
			break
		}
		go func() {
			//todo sth
			//step 3 回复数据
			fmt.Printf("addr: %v data: %v  count: %v\n", addr, string(data[:n]), n)
			_, err = listener.WriteToUDP([]byte("received success!"), addr)
			if err != nil {
				fmt.Printf("write failed, err: %v\n", err)
			}
		}()
	}
}
