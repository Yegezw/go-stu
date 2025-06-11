package main

import (
	"log"
	"net/http"
	"time"
)

// 简化
func test4() {
	// DefaultServerMux#HandleFunc -> DefaultServerMux#Handle
	// 向 DefaultServeMux#map[string]muxEntry 中增加对应的 handler 和路由规则
	http.HandleFunc("/hello", SayHello)

	// 实例化 Server
	// 调用 Server#ListenAndServe()
	// 调用 net.Listen("tcp", addr) 监听端口 -> for 循环 Accept 请求
	// 对每个请求实例化一个 Conn, 并开启一个 goroutine 为这个请求进行服务 go c.serve()
	// 读取每个请求的内容 w, err := c.readRequest()
	// 判断 handler 是否为空, 如果没有设置 handler, handler 就设置为 DefaultServeMux
	// 调用 handler 的 ServeHttp
	// 在这个例子中就进入到 DefaultServerMux#ServeHttp
	// 根据 request 选择 handler, 并且进入到这个 handler 的 ServeHTTP

	// 选择 handler
	// A 判断是否有路由能满足这个 request (循环遍历 ServerMux 的 muxEntry)
	// B 如果有路由满足, 调用这个路由 handler 的 ServeHttp
	// C 如果无路由满足, 调用 NotFoundHandler 的 ServeHttp
	_ = http.ListenAndServe(":8888", nil)
}

// 完整
func test5() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", SayHello)

	server := http.Server{
		Addr:        ":9999",
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func SayHello(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("Hello"))
}
