package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	//创建路由器
	mux := http.NewServeMux()
	// 设置路由规则,注册路由
	mux.HandleFunc("/byte", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(1 * time.Second)
		writer.Write([]byte("hello world, client!"))
	})

	//创建服务器
	server := &http.Server{
		Addr:         ":9090",
		WriteTimeout: time.Second * 15,
		Handler:      mux,
	}
	//监听端口并提供服务
	log.Println("start http server on :9090")
	log.Fatal(server.ListenAndServe())
}
