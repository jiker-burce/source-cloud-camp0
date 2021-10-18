package main

import (
	"fmt"
	"github.com/thinkeridea/go-extend/exnet"
	"log"
	"net/http"
	"os"
)

func healthZ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func full(w http.ResponseWriter, r *http.Request) {
	//1. 接收客户端 request，并将 request 中带的 header 写入 response header
	for name, values := range r.Header {
		w.Header().Set(name, values[0])
	}

	//2. 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	version := os.Getenv("VERSION")
	// 在~/.zshrc 中设置 export VERSION="v1.666"， source ~/.zshrc , 重启goland
	// 如果在可以在~/.bash_profile中设置了环境变量，可以在~/.zshrc中加入export ~/.bash_profile
	w.Header().Set("VERSION", version)

	//3. Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	returnCode := 204
	fmt.Println("IP:", exnet.ClientIP(r), "code:", returnCode, "version:", version)
	w.WriteHeader(returnCode)
}

func main() {
	http.HandleFunc("/healthz", healthZ) // 心跳健康检查
	http.HandleFunc("/full", full)       // 全量信息

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
