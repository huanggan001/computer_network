package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

var (
	host = []string{"127.0.0.1:2181"}
)

func main() {
	conn, _, err := zk.Connect(host, 5*time.Second)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	//增
	//if _, err := conn.Create("/test_tree", []byte("tree_content"), 0, zk.WorldACL(zk.PermAll)); err != nil {
	//	fmt.Println("create err: ", err)
	//}
	//查
	//nodeValue, dStat, err := conn.Get("/test_tree")
	//if err != nil {
	//	fmt.Println("get err : ", err)
	//	return
	//}
	//fmt.Println("nodeValue:", string(nodeValue))
	//fmt.Println(dStat)
	//改
	//if _, err := conn.Set("/test_tree", []byte("new_content"), dStat.Aversion); err != nil {
	//	fmt.Println("update err ", err)
	//}
	//删除
	//if err := conn.Delete("/test_tree", dStat.Version); err != nil {
	//	fmt.Println("delete err", err)
	//}
	//验证是否存在
	//hasNode, _, err := conn.Exists("/test_tree")
	//if err != nil {
	//	fmt.Println("exists err :", err)
	//}
	//fmt.Println("node exists :", hasNode)
	//增加
	//flag 参数的值为 0，这表示没有特殊的标志位被设置。这意味着新创建的节点将具有以下默认特性：
	//持久性：新节点将是持久的，即它将一直存在，直到显式删除。
	//顺序性：新节点将不带有顺序号，即节点名称不会包含任何序列号。
	//如果你想要设置特殊的节点标志，可以使用以下常量进行按位或运算：
	//zk.FlagEphemeral: 表示创建临时节点（如果会话结束，节点将被删除）。
	//zk.FlagSequence: 表示创建带有顺序号的节点。
	if _, err := conn.Create("/test_tree2", []byte("tree_content2"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("create err :", err)
	}
	//设置子节点
	if _, err := conn.Create("/test_tree2/subnode1", []byte("node_content1"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("create err :", err)
	}
	if _, err := conn.Create("/test_tree2/subnode2", []byte("node_content2"), 0, zk.WorldACL(zk.PermAll)); err != nil {
		fmt.Println("create err :", err)
	}
	//获取子节点列表
	childNodes, _, err := conn.Children("/test_tree2")
	if err != nil {
		fmt.Println("children err ", err)
	}
	fmt.Println("childNodes", childNodes)
}
