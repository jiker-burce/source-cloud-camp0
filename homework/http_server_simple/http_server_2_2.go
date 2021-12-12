package main

import (
	"20210221practice/src/cncamp/homework/http_server_simple/metrics"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/thinkeridea/go-extend/exnet"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
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

/**
启动时通过 -log_dir=log 指定目录，但是log目录必须存在
*/
func logs(w http.ResponseWriter, r *http.Request) {
	defer glog.Flush()
	//flag.Lookup("logtostderr").Value.Set("true")
	//currentPath, _ := os.Getwd()
	//flag.Lookup("log_dir").Value.Set(currentPath + "/log")
	flag.Parse()
	fmt.Print("log_dir:", flag.Lookup("log_dir").Value)

	now := time.Now().Format("2006-01-02 15:04:05")
	ip := exnet.ClientIP(r)
	statusCode := http.StatusOK
	w.Header().Add("statusCode", "200")
	path := r.RequestURI
	glog.V(2).Infof("%s\t%s\t%s\t%d", now, ip, path, statusCode)
}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func randTimeout(w http.ResponseWriter, r *http.Request){
	glog.V(4).Info("Random timeout")

	timer := metrics.NewTimer()
	defer timer.ObserveTotal()

	delayInt := randInt(10, 2000)

	time.Sleep(time.Millisecond * time.Duration(delayInt))
	glog.V(4).Infof("Timeout in %d ms", delayInt)
}

func main() {
	http.HandleFunc("/healthz", healthZ) // 心跳健康检查
	http.HandleFunc("/full", full)       // 全量信息
	http.HandleFunc("/logs", logs)       // 打印日志
	http.HandleFunc("/randTimeout", logs)       // 为 HTTPServer 添加 0-2 秒的随机延时

	err := http.ListenAndServe(":9090", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
