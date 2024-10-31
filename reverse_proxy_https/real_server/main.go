package main

import (
	"computer_network/reverse_proxy_https/testdata"
	"fmt"
	"golang.org/x/net/http2"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
证书签名生成方式:

//CA私钥
openssl genrsa -out yourdomain.com.key 4096
//生成证书签名请求(CSR)

	openssl req -sha512 -new \
		-subj "/C=CN/ST=Beijing/L=Beijing/O=example/OU=Personal/CN=yourdomain.com" \
		-key yourdomain.com.key \
		-out yourdomain.com.csr

//生成 x509 v3 扩展文件
cat > v3.ext <<-EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
extendedKeyUsage = serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.1=yourdomain.com
DNS.2=yourdomain
DNS.3=hostname
EOF

//使用该文件 v3.ext 为您的 Harbor 主机生成证书。

	openssl x509 -req -sha512 -days 3650 \
	    -extfile v3.ext \
	    -CA ca.crt -CAkey ca.key -CAcreateserial \
	    -in yourdomain.com.csr \
	    -out yourdomain.com.crt
*/
type RealServer struct {
	Addr string
}

func main() {
	rs1 := &RealServer{Addr: "127.0.0.1:3003"}
	rs1.Run()
	rs2 := &RealServer{Addr: "127.0.0.1:3004"}
	rs2.Run()
	//监听关闭信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func (r *RealServer) Run() {
	log.Println("Starting httpserver at " + r.Addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", r.HelloHandler)
	mux.HandleFunc("/base/error", r.ErrorHandler)
	server := &http.Server{
		Addr:         r.Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	go func() {
		http2.ConfigureServer(server, &http2.Server{})
		log.Fatal(server.ListenAndServeTLS(testdata.Path("huanggan.com.crt"), testdata.Path("huanggan.com.key")))
	}()
}
func (r *RealServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
	upath := fmt.Sprintf("http://%s%s---huanggan.com\n", r.Addr, req.URL.Path)
	io.WriteString(w, upath)
}
func (r *RealServer) ErrorHandler(w http.ResponseWriter, req *http.Request) {
	upath := "error handler"
	w.WriteHeader(500)
	io.WriteString(w, upath)
}
