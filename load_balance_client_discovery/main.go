package main

import (
	"computer_network/load_balance"
	"computer_network/middleware"
	proxy2 "computer_network/proxy"
	"context"

	"log"
	"net/http"
)

var addr = "127.0.0.1:2002"

func main() {
	mConf, err := load_balance.NewLoadBalanceCheckConf("http://%s/base", map[string]string{"127.0.0.1:2003": "20"})
	if err != nil {
		panic(err)
	}
	rb := load_balance.LoadBalanceFactorWithConf(load_balance.LbWeightRoundRobin, mConf)
	proxy := proxy2.NewLoadBalanceReverseProxy(&middleware.SliceRouterContext{Ctx: context.Background()}, rb)
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
