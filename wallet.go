package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

type Response struct {
    Data       interface{} `json:"data,omitempty"`
    Status     int         `json:"status"`
    StatusInfo *StatusInfo `json:"statusInfo,omitempty"`
}

type StatusInfo struct {
    Message string      `json:"message"`
    Detail  interface{} `json:"detail,omitempty"`
}

// UserSettings 包含用户配置信息
type UserSettings struct {
    UserID       int     `json:"userID"`
    Username     string  `json:"username"`
    Email        string  `json:"email"`
    AccountLevel string  `json:"accountLevel"`
    JoinDate     string  `json:"joinDate"`
    Balance      float64 `json:"balance"`
}

// UserSettingsRequest 表示设置API的请求载荷
type UserSettingsRequest struct {
    UserID string `json:"user_id"`
}

// 定义钱包资产相关结构体

// WalletAsset 表示用户钱包中的一种资产
type WalletAsset struct {
    Name           string  `json:"name"`
    Symbol         string  `json:"symbol"`
    Market24h      string  `json:"24h"`
    CurQuantity    float64 `json:"CurQuantity,"`
    CurValue       float64 `json:"CurValue,"`
    OnOrderQuantity float64 `json:"onOrderQuantity"`
    OnOrderValue   float64 `json:"onOrderValue,"`
}

// WalletInfo 包含用户钱包的资产信息
type WalletInfo struct {
    UserID   int           `json:"userID"`
    Balance  float64       `json:"balance"`
    Assets   []WalletAsset `json:"assets"`
}

// WalletInfoRequest 表示钱包API的请求载荷
type WalletInfoRequest struct {
    UserID string `json:"user_id"`
}

func main() {
    r := mux.NewRouter()
    
    // 用户接口路由
    userRouter := r.PathPrefix("/api/user").Subrouter()
    
    // 设置路由
    userRouter.HandleFunc("/wallet/settings", GetUserSettings).Methods("POST")
    userRouter.HandleFunc("/wallet/assets", GetWalletAssetInfo).Methods("POST")
    
    log.Fatal(http.ListenAndServe(":8080", r))
}

// GetUserSettings 获取用户配置信息
func GetUserSettings(w http.ResponseWriter, r *http.Request) {
    var req UserSettingsRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        RespondWithError(w, 1400, "无效的请求格式", map[string]string{"exception": "BadRequestException"})
        return
    }
    
    // 验证用户ID
    if req.UserID == "" {
        RespondWithError(w, 1400, "用户ID不能为空", map[string]string{"exception": "ValidationException"})
        return
    }
    
    // 验证用户是否存在
    exists, err := ValidateUserExists(req.UserID)
    if err != nil {
        RespondWithError(w, 1500, "系统错误", map[string]string{"exception": "InternalServerException"})
        log.Printf("验证用户失败: %v", err)
        return
    }
    
    if !exists {
        RespondWithError(w, 1001, "用户id错误，不存在", map[string]string{"exception": "InvalidCredentialsException"})
        return
    }
    
    // 从数据库获取用户设置信息
    userSettings, err := GetUserSettingsFromDB(req.UserID)
    if err != nil {
        RespondWithError(w, 1500, "获取用户设置信息失败", map[string]string{"exception": "InternalServerException"})
        log.Printf("获取用户设置失败: %v", err)
        return
    }
    
    // 返回用户设置信息
    RespondWithSuccess(w, userSettings)
}

// GetWalletAssetInfo 获取用户钱包资产信息
func GetWalletAssetInfo(w http.ResponseWriter, r *http.Request) {
    var req WalletInfoRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        RespondWithError(w, 1400, "无效的请求格式", map[string]string{"exception": "BadRequestException"})
        return
    }
    
    // 验证用户ID
    if req.UserID == "" {
        RespondWithError(w, 1400, "用户ID不能为空", map[string]string{"exception": "ValidationException"})
        return
    }
    
    // 验证用户是否存在
    exists, err := ValidateUserExists(req.UserID)
    if err != nil {
        RespondWithError(w, 1500, "系统错误", map[string]string{"exception": "InternalServerException"})
        log.Printf("验证用户失败: %v", err)
        return
    }
    
    if !exists {
        RespondWithError(w, 1001, "用户id错误，不存在", map[string]string{"exception": "InvalidCredentialsException"})
        return
    }
    
    // 从数据库获取用户钱包资产信息
    walletInfo, err := GetUserWalletAssets(req.UserID)
    if err != nil {
        RespondWithError(w, 1500, "获取钱包资产信息失败", map[string]string{"exception": "InternalServerException"})
        log.Printf("获取钱包资产失败: %v", err)
        return
    }
    
    // 返回用户钱包资产信息
    RespondWithSuccess(w, walletInfo)
}

// RespondWithSuccess 处理成功响应
func RespondWithSuccess(w http.ResponseWriter, data interface{}) {
    response := Response{
        Data:   data,
        Status: 0,
    }
    
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// RespondWithError 处理错误响应
func RespondWithError(w http.ResponseWriter, statusCode int, message string, detail interface{}) {
    response := Response{
        Status: statusCode,
        StatusInfo: &StatusInfo{
            Message: message,
            Detail:  detail,
        },
    }
    
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}