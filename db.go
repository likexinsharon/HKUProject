package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// User 用户模型
// Removed duplicate declaration. The User struct is defined in auth.go.

var db *sql.DB

func connectDB() error {
	var err error
	// 修正数据库连接字符串格式
	db, err = sql.Open("mysql", "root:hkuproject2025!@tcp(rm-bp1j0x5f9je2tr4fh.mysql.rds.aliyuncs.com:3306)/hkuproject")
	if err != nil {
		return err
	}
	// See "Important settings" section
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("db conncetion succ")

	return nil
}

// 检查用户名是否已存在
func isUsernameExists(username string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	return count > 0, err
}

// 检查邮箱是否已存在
func isEmailExists(email string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	return count > 0, err
}

// 根据用户名获取用户
func getUserByUsername(username string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password, email, account_level FROM users WHERE username = ?",
		username).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.AccountLevel)
	return user, err
}

// 根据邮箱获取用户
func getUserByEmail(email string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password, email, account_level FROM users WHERE email = ?",
		email).Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.AccountLevel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, fmt.Errorf("email not registered")
		}
		return User{}, err
	}
	return user, nil
}

// 创建用户
func createUser(user User) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO users (username, password, email) VALUES (?, ?, ?)",
		user.Username, user.Password, user.Email,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func isUsernameExistsMock(username string) (bool, error) {
	// 假设我们有一个用户列表
	existingUsers := []string{"admin", "guest", "user1"}
	for _, user := range existingUsers {
		if user == username {
			return true, nil
		}
	}
	return false, nil
}
