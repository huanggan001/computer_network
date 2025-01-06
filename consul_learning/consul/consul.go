package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net"
)

type Consul struct {
	Client *api.Client
}

// NewConsul 连接至consul服务，返回一个consul对象
func NewConsul(addr string) (*Consul, error) {
	cfg := api.DefaultConfig()
	cfg.Address = addr
	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Consul{Client: client}, nil
}

// GetOutboundIP 获取本机出口ip
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

// RegisterService 将grpc服务注册到consul
func (c *Consul) RegisterService(serviceName string, ip string, port int) error {

	//添加健康检查
	check := &api.AgentServiceCheck{
		GRPC:     fmt.Sprintf("%s:%d", ip, port), //这里一定是外部可以访问的地址(grpc服务地址)
		Timeout:  "5s",                           //超时时间
		Interval: "10s",                          //检查间隔,包括等待服务响应的时间,不是等待心跳返回后，才重新计时
		//指定时间后自动注销不健康的服务节点
		//最小超时时间为1分钟，收获不健康服务的进程每30秒运行一次，因此触发注销的时间可能略长于配置的超时时间。
		DeregisterCriticalServiceAfter: "1m",
	}

	srv := &api.AgentServiceRegistration{
		ID:      fmt.Sprintf("%s-%s-%d", serviceName, ip, port), //服务唯一id
		Name:    serviceName,                                    //服务名称
		Tags:    []string{"grpc", "consul_learning"},            //为服务打标签
		Address: ip,
		Port:    port,
		Check:   check,
	}
	return c.Client.Agent().ServiceRegister(srv)
}
