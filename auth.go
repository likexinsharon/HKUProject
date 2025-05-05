package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// User 用户模型
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"-"` // 不在JSON中返回密码
	Email        string `json:"email"`
	AccountLevel string `json:"accountLevel"`
}

// JWT配置
const (
	jwtSecret        = "hku-project-secret-key-2025" // 实际应用中应放在环境变量中
	tokenExpireTime  = 24 * time.Hour                // 默认token过期时间
	longTokenExpTime = 7 * 24 * time.Hour            // 记住登录的token过期时间
)

// Claims JWT声明结构
type Claims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

// 生成密码哈希
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// 验证密码
func verifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// 生成JWT令牌
func generateToken(userID int, rememberMe bool) (string, int64, error) {
	// 设置过期时间
	expirationTime := time.Now().Add(tokenExpireTime)
	if rememberMe {
		expirationTime = time.Now().Add(longTokenExpTime)
	}

	// 创建声明
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// 创建令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expirationTime.Unix() - time.Now().Unix(), nil
}

// 验证JWT令牌
func verifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("意外的签名方法: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	// 处理解析错误
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("令牌已过期")
		}
		return nil, errors.New("令牌无效")
	}

	// 验证令牌有效性
	if !token.Valid {
		return nil, errors.New("令牌无效")
	}

	return claims, nil
}

// 修正 extractTokenFromRequest 函数，避免 body 被多次读取的问题
func extractTokenFromRequest(r *http.Request) string {
	// 从Authorization头获取
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	// 从查询参数获取
	token := r.URL.Query().Get("token")
	if token != "" {
		return token
	}

	// 从请求体获取（不再尝试读取请求体，避免干扰后续处理）
	return ""
}

// 生成随机密码重置令牌
func generateResetToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// 处理登录请求
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, 1001, "方法不允许", "MethodNotAllowedException")
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, 1001, "无效的请求格式", "InvalidRequestException")
		return
	}

	// 获取用户信息
	user, err := getUserByUsername(req.Username)
	if err != nil {
		sendErrorResponse(w, 1001, "用户名或密码错误", "InvalidCredentialsException")
		return
	}

	// 验证密码
	if !verifyPassword(user.Password, req.Password) {
		sendErrorResponse(w, 1001, "用户名或密码错误", "InvalidCredentialsException")
		return
	}

	// 生成JWT令牌
	token, _, err := generateToken(user.ID, req.RememberMe)
	if err != nil {
		sendErrorResponse(w, 1001, "生成令牌失败", "TokenGenerationException")
		return
	}

	// 准备响应
	response := ApiResponse{
		Status: 0,
		Data: map[string]interface{}{
			"token": token,
			"user": map[string]interface{}{
				"id":           user.ID,
				"username":     user.Username,
				"email":        user.Email,
				"accountLevel": user.AccountLevel,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// 处理注册请求
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, 1002, "方法不允许", "MethodNotAllowedException")
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, 1002, "无效的请求格式", "InvalidRequestException")
		return
	}

	// 验证必填字段
	if req.Username == "" || req.Password == "" || req.Email == "" {
		sendErrorResponse(w, 1002, "请填写所有必填字段", "MissingRequiredFieldException")
		return
	}

	// 检查是否同意条款
	if !req.AgreeTerms {
		sendErrorResponse(w, 1002, "必须同意服务条款", "TermsNotAgreedException")
		return
	}

	// 检查用户名是否已存在
	exists, err := isUsernameExists(req.Username)
	if err != nil {
		sendErrorResponse(w, 1002, "验证用户名失败", "DatabaseException")
		return
	}
	if exists {
		sendErrorResponse(w, 1002, "用户名已存在", "DuplicateUsernameException")
		return
	}

	// 检查邮箱是否已存在
	exists, err = isEmailExists(req.Email)
	if err != nil {
		sendErrorResponse(w, 1002, "验证邮箱失败", "DatabaseException")
		return
	}
	if exists {
		sendErrorResponse(w, 1002, "邮箱已被注册", "DuplicateEmailException")
		return
	}

	// 哈希密码
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		sendErrorResponse(w, 1002, "密码处理失败", "PasswordHashingException")
		return
	}

	// 创建用户
	user := User{
		Username:     req.Username,
		Password:     hashedPassword,
		Email:        req.Email,
		AccountLevel: "standard", // 默认账户级别
	}

	userID, err := createUser(user)
	if err != nil {
		sendErrorResponse(w, 1002, "创建用户失败", "DatabaseException")
		return
	}

	// 准备响应
	response := ApiResponse{
		Status: 0,
		Data: map[string]interface{}{
			"userId":               userID,
			"username":             req.Username,
			"verificationRequired": true,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// 处理重置密码请求
func resetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, 1003, "方法不允许", "MethodNotAllowedException")
		return
	}

	var req ResetPasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, 1003, "无效的请求格式", "InvalidRequestException")
		return
	}

	// 检查邮箱是否存在
	user, err := getUserByEmail(req.Email)
	if err != nil {
		sendErrorResponse(w, 1003, "该邮箱未注册", "EmailNotFoundException")
		return
	}

	// 生成密码重置令牌
	resetToken, err := generateResetToken()
	if err != nil {
		sendErrorResponse(w, 1003, "生成重置令牌失败", "TokenGenerationException")
		return
	}

	// TODO: 在实际应用中，应该将令牌存储到数据库并发送重置邮件
	// 此处为简化处理，仅返回成功响应
	fmt.Printf("为用户 %s 生成了密码重置令牌: %s\n", user.Username, resetToken)

	// 准备响应
	response := ApiResponse{
		Status: 0,
		Data: map[string]interface{}{
			"emailSent":         true,
			"expirationMinutes": 30,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// 处理登出请求
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, 1001, "方法不允许", "MethodNotAllowedException")
		return
	}

	// 获取令牌
	// 获取令牌并在实际应用中将其加入黑名单
	_ = extractTokenFromRequest(r)

	// TODO: 将令牌添加到黑名单中

	// 准备响应
	response := ApiResponse{
		Status: 0,
		Data: map[string]interface{}{
			"logoutSuccess": true,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// 处理验证令牌请求
func verifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, 1004, "方法不允许", "MethodNotAllowedException")
		return
	}

	var req VerifyTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendErrorResponse(w, 1004, "无效的请求格式", "InvalidRequestException")
		return
	}

	// 验证令牌
	claims, err := verifyToken(req.Token)
	if err != nil {
		sendErrorResponse(w, 1004, err.Error(), "TokenExpiredException")
		return
	}

	// 计算过期时间（秒）
	expiresIn := claims.ExpiresAt.Unix() - time.Now().Unix()

	// 准备响应
	response := ApiResponse{
		Status: 0,
		Data: map[string]interface{}{
			"valid":     true,
			"expiresIn": expiresIn,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// 辅助函数，发送错误响应
func sendErrorResponse(w http.ResponseWriter, status int, message string, exception string) {
	response := ApiResponse{
		Status: status,
		StatusInfo: &StatusInfo{
			Message: message,
			Detail: map[string]interface{}{
				"exception": exception,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// 确保初始化测试数据函数在这里定义
func initUserTestData() {
	// 检查是否已有测试用户
	exists, _ := isUsernameExists("testuser")
	if !exists {
		hashedPassword, _ := hashPassword("password123")
		user := User{
			Username:     "testuser",
			Password:     hashedPassword,
			Email:        "test@example.com",
			AccountLevel: "standard",
		}
		createUser(user)
		fmt.Println("创建测试用户: testuser")
	}
}
