package main

import (
	"computer_network/zookeeper/public"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var addr = "127.0.0.1:2181"

func main() {
	//获取zk节点列表
	zkManager := public.NewZkManager([]string{"127.0.0.1:2181"})
	zkManager.GetConnect()
	defer zkManager.Close()

	zlist, err := zkManager.GetServerListByPath("/real_server")
	fmt.Println("server nodes:")
	fmt.Println(zlist)
	if err != nil {
		log.Println(err)
	}

	//动态监听节点数变化
	chanList, chanErr := zkManager.WatchServerListByPath("/real_server")
	go func() {
		for {
			select {
			case changeErr := <-chanErr:
				fmt.Println("changeErr")
				fmt.Println(changeErr)
			case changedList := <-chanList:
				fmt.Println("watch node changed")
				fmt.Println(changedList)
			}
		}
	}()

	//动态监听节点内容
	//dataList, chanErr := zkManager.WatchServerListByPath("/real_server")
	//dataList, chanErr := zkManager.WatchPathData("/real_server")
	//go func() {
	//	for {
	//		select {
	//		case changeErr := <-chanErr:
	//			fmt.Println("changeErr")
	//			fmt.Println(changeErr)
	//		case changedList := <-dataList:
	//			fmt.Println("watch data changed:")
	//			fmt.Println(string(changedList))
	//		}
	//	}
	//}()

	//关闭信号监听
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
