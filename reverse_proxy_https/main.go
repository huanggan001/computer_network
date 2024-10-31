package main

import (
	"computer_network/reverse_proxy_https/public"
	"computer_network/reverse_proxy_https/testdata"
	"log"
	"net/http"
	"net/url"
)

var addr = "127.0.0.1:3002"

func main() {
	rs1 := "https://huanggan.com:3003"
	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}
	urls := []*url.URL{url1}
	proxy := public.NewMultipleHostsReverseProxy(urls)
	log.Println("Starting httpserver at " + addr)

	log.Fatal(http.ListenAndServeTLS(addr, testdata.Path("huanggan.com.crt"), testdata.Path("huanggan.com.key"), proxy))
}
