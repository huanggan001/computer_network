package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	consulAddress = "127.0.0.1:8500"
	localIp       = "10.60.82.38"
	localPort     = 81
)

func consulRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "10"
	registration.Name = "service10"
	registration.Port = localPort
	registration.Tags = []string{"testService"}
	registration.Address = localIp

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println("Service registration error: ", err)
	}
}

// 注销服务
func DeRegister(serviceID string) {
	client, _ := consulapi.NewClient(&consulapi.Config{Address: "127.0.0.1:8500"})
	_, _ = client.Catalog().Deregister(&consulapi.CatalogDeregistration{
		//consul服务器的节点名称
		Node:      "consul1",
		Address:   localIp,
		ServiceID: serviceID,
	}, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("you are visiting health check api"))
}

func main() {
	consulRegister()
	//defer DeRegister("10")
	// 捕获退出信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 在一个 goroutine 中监听信号
	go func() {
		<-sigs
		DeRegister("10")
		os.Exit(0)
	}()
	//定义一个http接口
	//一定要处理ip:port/这个请求，不然即使服务运行的好好的，健康检查也是无法通过的。
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe("0.0.0.0:81", nil)
	if err != nil {
		fmt.Println("error: ", err.Error())
	}
}
