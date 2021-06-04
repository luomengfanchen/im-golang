package main

import (
	"fmt"
	"im/utils"
	"net/http"
)

func init() {
	// 加载json配置
	utils.LoadConfig("config.json")

	// 服务启动输出
	fmt.Println("im server syestem start.")
	fmt.Println("at: http://", utils.Config.Address)
}

func main() {
	// 创建多路复用器
	mux := http.NewServeMux()

	// 静态文件服务
	file := http.Dir(utils.Config.Static)
	mux.Handle("/", http.FileServer(file))

	// 配置http服务
	server := &http.Server{
		Addr:    utils.Config.Address,
		Handler: mux,
	}

	// 开启监听服务
	server.ListenAndServe()
}
