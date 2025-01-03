package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	//创建连接池
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, //连接超时
			KeepAlive: 30 * time.Second, // 探活时间
		}).DialContext,
		MaxIdleConns:          100,              //最大空闲连接
		IdleConnTimeout:       90 * time.Second, //空闲超时时间
		TLSHandshakeTimeout:   10 * time.Second, //tls握手超时时间
		ExpectContinueTimeout: 1 * time.Second,  //100-continue状态码超时时间
	}
	//创建客户端
	client := &http.Client{
		Timeout:   30 * time.Second, //请求超时时间
		Transport: transport,
	}

	//请求数据
	resp, err := client.Get("http://127.0.0.1:9090/byte")
	if err != nil {
		//使用panic，而不是return
		//是因为panic在终止程序前，会执行defer的任务
		panic(err)
		//return
	}
	defer resp.Body.Close()
	//读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
