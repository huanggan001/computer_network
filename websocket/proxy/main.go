package main

import (
	"computer_network/middleware"
	proxy2 "computer_network/proxy"
	"context"
	"log"
	"net/http"
	"net/url"
)

var addr = "127.0.0.1:8081"

func main() {
	addr2, _ := url.Parse("http://127.0.0.1:8080")

	proxy := proxy2.NewMultipleHostsReverseProxy(&middleware.SliceRouterContext{Ctx: context.Background()}, []*url.URL{addr2})
	log.Println("Starting httpserver at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
