package main

import (
	"fmt"
	"im/service"
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

	// 登录认证
	mux.HandleFunc("/api/login", service.Login)

	// 好友列表
	mux.HandleFunc("/api/friend", service.Friend)

	// 获取聊天记录
	mux.HandleFunc("/api/chatlist", service.ChatList)

	// 发送聊天记录
	mux.HandleFunc("/api/chat", service.Chat)

	// 配置http服务
	server := &http.Server{
		Addr:    utils.Config.Address,
		Handler: mux,
	}

	// 开启监听服务
	server.ListenAndServe()
}
