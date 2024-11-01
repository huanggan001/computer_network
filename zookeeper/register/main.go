package main

import (
	"computer_network/zookeeper/public"
	"fmt"
	"time"
)

func main() {
	zkManager := public.NewZkManager([]string{"127.0.0.1:2181"})
	zkManager.GetConnect()
	defer zkManager.Close()
	i := 0
	for {
		zkManager.SetPathData("/real_server", []byte(fmt.Sprint(i)), 0)
		//zkManager.RegistServerPath("/real_server", fmt.Sprint(i))
		fmt.Println("update data -> ", i)
		time.Sleep(5 * time.Second)
		i++
	}
}
