package main

// API请求结构
type LoginRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
}

type RegisterRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	AgreeTerms bool   `json:"agreeTerms"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

// API响应结构
type ApiResponse struct {
	Status     int                    `json:"status"`
	Data       map[string]interface{} `json:"data,omitempty"`
	StatusInfo *StatusInfo            `json:"statusInfo,omitempty"`
}
