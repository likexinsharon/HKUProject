package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func GetUserSettings(w http.ResponseWriter, r *http.Request) {
	// 用户设置处理逻辑
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("用户设置已获取"))
}

func GetWalletAssetInfo(w http.ResponseWriter, r *http.Request) {
	// 钱包资产信息处理逻辑
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("钱包资产信息已获取"))
}

func main() {
	// 连接数据库
	fmt.Println("正在连接数据库...")
	err := connectDB()
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	fmt.Println("数据库连接成功")

	// 初始化测试数据
	initTestData()

	r := mux.NewRouter()

	// 注册原有路由
	r.HandleFunc("/api/hello", loggingMiddleware(corsMiddleware(helloHandler)))

	// 用户接口路由
	userRouter := r.PathPrefix("/api/user").Subrouter()

	// 注册用户认证相关路由
	userRouter.HandleFunc("/login", loggingMiddleware(corsMiddleware(loginHandler))).Methods("POST")
	userRouter.HandleFunc("/register", loggingMiddleware(corsMiddleware(registerHandler))).Methods("POST")
	userRouter.HandleFunc("/reset-password", loggingMiddleware(corsMiddleware(resetPasswordHandler))).Methods("POST")
	userRouter.HandleFunc("/logout", loggingMiddleware(corsMiddleware(logoutHandler))).Methods("POST")
	userRouter.HandleFunc("/verify-token", loggingMiddleware(corsMiddleware(verifyTokenHandler))).Methods("POST")
	userRouter.HandleFunc("/wallet/settings", GetUserSettings).Methods("POST")
	userRouter.HandleFunc("/wallet/assets", GetWalletAssetInfo).Methods("POST")

	fmt.Println("服务器正在运行于 :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
