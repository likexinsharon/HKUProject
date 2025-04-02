package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// 定义请求/响应用的结构体
type Message struct {
	Name string `json:"name"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 只接受 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 解析请求体中的 JSON
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// 构造响应
	response := map[string]string{
		"message": fmt.Sprintf("Hello, %s!", msg.Name),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 为所有请求添加CORS头的中间件
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

// 记录请求信息的中间件
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("[%s] %s %s\n", start.Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		next(w, r)
		fmt.Printf("[%s] %s %s 完成，耗时: %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.Method, r.URL.Path,
			time.Since(start))
	}
}

func main() {
	// 连接数据库
	println("start")
	err := connectDB()
	println("end")
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	fmt.Println("数据库连接成功")

	// 初始化测试数据
	initTestData()

	// 注册原有路由
	http.HandleFunc("/api/hello", loggingMiddleware(corsMiddleware(helloHandler)))

	// 注册用户认证相关路由
	http.HandleFunc("/api/user/login", loggingMiddleware(corsMiddleware(loginHandler)))
	http.HandleFunc("/api/user/register", loggingMiddleware(corsMiddleware(registerHandler)))
	http.HandleFunc("/api/user/reset-password", loggingMiddleware(corsMiddleware(resetPasswordHandler)))
	http.HandleFunc("/api/user/logout", loggingMiddleware(corsMiddleware(logoutHandler)))
	http.HandleFunc("/api/user/verify-token", loggingMiddleware(corsMiddleware(verifyTokenHandler)))

	fmt.Println("服务器正在运行于 127.0.0.1:80")
	log.Fatal(http.ListenAndServe("127.0.0.1:80", nil))
}
