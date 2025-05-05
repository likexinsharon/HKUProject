package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
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

// GetUserSettingsFromDB 从数据库获取用户设置信息
func GetUserSettingsFromDB(userID string) (UserSettings, error) {
	var settings UserSettings
	var joinDate string

	// 连接检查
	if db == nil {
		return settings, fmt.Errorf("数据库连接未初始化")
	}

	// 查询用户设置信息
	query := `
		SELECT us.user_id, us.username, us.email, us.account_level, us.join_date 
		FROM user_settings us
		JOIN users u ON us.user_id = u.id
		WHERE u.id = ?`

	err := db.QueryRow(query, userID).Scan(
		&settings.UserID,
		&settings.Username,
		&settings.Email,
		&settings.AccountLevel,
		&joinDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return settings, fmt.Errorf("用户ID不存在")
		}
		return settings, fmt.Errorf("查询用户设置失败: %w", err)
	}

	// 格式化日期
	settings.JoinDate = joinDate

	// 计算用户余额
	balance, err := GetUserTotalBalance(userID)
	if err != nil {
		return settings, fmt.Errorf("计算用户余额失败: %w", err)
	}

	settings.Balance = balance

	return settings, nil
}

// GetUserTotalBalance 计算用户所有资产的总价值
func GetUserTotalBalance(userID string) (float64, error) {
	var totalBalance float64
	// todo ：多个历史模拟交易（需要做）
	query := `
		SELECT SUM(ua.quantity * ap.price) + SUM(ua.on_order_quantity * ap.price) as total_balance
		FROM user_assets ua
		JOIN asset_prices ap ON ua.symbol = ap.symbol
		WHERE ua.user_id = ?`

	err := db.QueryRow(query, userID).Scan(&totalBalance)
	if err != nil {
		return 0, err
	}

	return totalBalance, nil
}

// GetUserWalletAssets 获取用户钱包的资产信息
func GetUserWalletAssets(userID string) (WalletInfo, error) {
	var walletInfo WalletInfo
	var err error

	// 连接检查
	if db == nil {
		return walletInfo, fmt.Errorf("数据库连接未初始化")
	}

	// 设置用户ID
	id, err := strconv.Atoi(userID)
	if err != nil {
		return walletInfo, fmt.Errorf("无效的用户ID格式")
	}
	walletInfo.UserID = id

	// 获取总余额
	balance, err := GetUserTotalBalance(userID)
	if err != nil {
		return walletInfo, fmt.Errorf("计算用户余额失败: %w", err)
	}
	walletInfo.Balance = balance

	// TODO: 查询用户资产列表  【表结构待确认】
	query := `
		SELECT ua.asset_name, 
		       ua.symbol, 
		       CONCAT(IF(ap.change_24h >= 0, '+', ''), 
		              CAST(ap.change_24h AS CHAR), 
		              '%') AS market_24h,
		       ua.quantity, 
		       ua.quantity * ap.price AS current_value,
		       ua.on_order_quantity, 
		       ua.on_order_quantity * ap.price AS on_order_value
		FROM user_assets ua
		FORCE INDEX (idx_user_id)
		INNER JOIN asset_prices ap 
		    ON ua.symbol = ap.symbol AND ap.symbol IN (
		        SELECT DISTINCT symbol FROM user_assets 
		        WHERE user_id = ?
		    )
		WHERE ua.user_id = ?
		ORDER BY current_value DESC
		LIMIT 50;`

	rows, err := db.Query(query, userID, userID)
	if err != nil {
		return walletInfo, fmt.Errorf("查询用户资产失败: %w", err)
	}
	defer rows.Close()

	// 处理查询结果
	var assets []WalletAsset
	for rows.Next() {
		var asset WalletAsset
		err := rows.Scan(
			&asset.Name,
			&asset.Symbol,
			&asset.Market24h,
			&asset.CurQuantity,
			&asset.CurValue,
			&asset.OnOrderQuantity,
			&asset.OnOrderValue,
		)
		if err != nil {
			return walletInfo, fmt.Errorf("处理资产数据失败: %w", err)
		}
		assets = append(assets, asset)
	}

	if err := rows.Err(); err != nil {
		return walletInfo, fmt.Errorf("迭代资产数据失败: %w", err)
	}

	walletInfo.Assets = assets

	return walletInfo, nil
}

// 初始化钱包测试数据
func initWalletTestData() error {
	// 添加测试用户设置
	userSettingsQuery := `
		INSERT IGNORE INTO user_settings (user_id, username, email, account_level, join_date)
		VALUES 
			(12345, 'testuser', 'test@example.com', 'standard', '2020-10-20'),
			(12345678, 'premium', 'premium@example.com', 'vip', '2019-05-15')`

	// 添加测试资产价格
	pricesQuery := `
		INSERT IGNORE INTO asset_prices (symbol, price, change_24h)
		VALUES 
			('BTC', 50000, 2.05),
			('ETH', 3000, -2.05),
			('BNB', 350, 1.23),
			('SOL', 80, 5.67)`

	// 添加测试用户资产
	assetsQuery := `
		INSERT IGNORE INTO user_assets (user_id, asset_name, symbol, quantity, on_order_quantity)
		VALUES 
			(12345, 'Bitcoin', 'BTC', 2, 0.00674),
			(12345, 'Ethereum', 'ETH', 33.33, 1.179),
			(12345678, 'Bitcoin', 'BTC', 5.5, 0.1),
			(12345678, 'Ethereum', 'ETH', 75, 2.5),
			(12345678, 'Binance Coin', 'BNB', 100, 10)`

	// 执行SQL
	if _, err := db.Exec(userSettingsQuery); err != nil {
		return fmt.Errorf("初始化用户设置数据失败: %w", err)
	}

	if _, err := db.Exec(pricesQuery); err != nil {
		return fmt.Errorf("初始化资产价格数据失败: %w", err)
	}

	if _, err := db.Exec(assetsQuery); err != nil {
		return fmt.Errorf("初始化用户资产数据失败: %w", err)
	}

	return nil
}

// ValidateUserExists 验证用户ID是否存在
func ValidateUserExists(userID string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)"
	err := db.QueryRow(query, userID).Scan(&exists)
	return exists, err
}

// initTestData 函数更新
func initTestData() {
	// 现有测试数据初始化代码...

	// 初始化钱包测试数据
	if err := initWalletTestData(); err != nil {
		log.Printf("初始化钱包测试数据失败: %v", err)
	} else {
		fmt.Println("钱包测试数据初始化成功")
	}
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
