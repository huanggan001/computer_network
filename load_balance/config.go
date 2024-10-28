package load_balance

import (
	"computer_network/zookeeper/public"
	"fmt"
)

type Oberver interface {
	Update()
}

// 配置主题
type LoadBalanceConf interface {
	Attach(o Oberver)
	GetConf() []string
	WatchConf()
	UpdateConf(conf []string)
}

type LoadBalanceZkConf struct {
	observers    []Oberver
	path         string
	zkHosts      []string          //zookeeper集群列表
	confIpWeight map[string]string //根据ip获取权重
	activeList   []string          //活跃的服务器列表
	format       string
}

// 添加观察者
func (s *LoadBalanceZkConf) Attach(o Oberver) {
	s.observers = append(s.observers, o)
}

func (s *LoadBalanceZkConf) NotifyAllObervers() {
	for _, obs := range s.observers {
		obs.Update()
	}
}

func (s *LoadBalanceZkConf) GetConf() []string {
	confList := []string{}
	for _, ip := range s.activeList {
		weight, ok := s.confIpWeight[ip]
		if !ok {
			weight = "50" //默认weight
		}
		confList = append(confList, fmt.Sprintf(s.format, ip)+","+weight)
	}
	return confList
}

// 监听配置
func (s *LoadBalanceZkConf) WatchConf() {
	zkManager := public.NewZkManager(s.zkHosts)
	zkManager.GetConnect()
	chanList, chanErr := zkManager.WatchServerListByPath(s.path)
	go func() {
		defer zkManager.Close()
		for {
			select {
			case changeErr := <-chanErr:
				fmt.Println("changeErr", changeErr)
			case changedList := <-chanList:
				fmt.Println("watch node changed")
				s.UpdateConf(changedList)
			}
		}
	}()
}

// 更新配置，通知监听者也更新
func (s *LoadBalanceZkConf) UpdateConf(conf []string) {
	s.activeList = conf
	for _, obs := range s.observers {
		obs.Update()
	}
}

func NewLoadBalanceZkConf(format, path string, zkHosts []string, conf map[string]string) (*LoadBalanceZkConf, error) {
	zkManager := public.NewZkManager(zkHosts)
	zkManager.GetConnect()
	defer zkManager.Close()
	zlist, err := zkManager.GetServerListByPath(path)
	if err != nil {
		return nil, err
	}
	mConf := &LoadBalanceZkConf{format: format, activeList: zlist, confIpWeight: conf, zkHosts: zkHosts, path: path}
	mConf.WatchConf()
	return mConf, nil
}

type LoadBalanceObserver struct {
	ModuleConf *LoadBalanceZkConf
}

func (l *LoadBalanceObserver) Update() {
	fmt.Println("Update get conf:", l.ModuleConf.GetConf())
}

func NewLoadBalanceObserver(conf *LoadBalanceZkConf) *LoadBalanceObserver {
	return &LoadBalanceObserver{
		ModuleConf: conf,
	}
}
