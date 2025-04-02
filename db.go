package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// ...

func connectDB() error {
	db, err := sql.Open("mysql", "root:hkuproject2025!@/rm-bp1j0x5f9je2tr4fh.mysql.rds.aliyuncs.com")
	if err != nil {
		return err
	}
	// See "Important settings" section
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return nil
}
