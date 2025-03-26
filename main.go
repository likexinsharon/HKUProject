package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
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

func main() {
    http.HandleFunc("/api/hello", helloHandler)

    fmt.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
