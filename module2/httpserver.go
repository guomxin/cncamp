package main

import (
	"fmt"
	"log"
	"module2/utils"
	"net/http"
	"os"

	_ "github.com/spf13/cobra"
)

func main() {
	// 设置路由处理函数
	http.HandleFunc("/healthz", healthzHandler)

	// 启动HTTP服务器
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// 将请求头信息写入响应头
	utils.WriteReqHeadersToResponse(w, r)

	// 读取环境变量中的VERSION配置，并写入响应头
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Set("Version", version)
	}

	// 记录访问日志
	fmt.Printf("Client IP: %s, HTTP status code: %d\n", r.RemoteAddr, http.StatusOK)

	// 返回响应
	fmt.Fprintf(w, "Hello, World!")

	// 返回200状态码
	w.WriteHeader(http.StatusOK)
}
