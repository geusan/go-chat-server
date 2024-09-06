package rdb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: make environment variable
var dsn = "root:localhost@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"

func OpenTestDB(Conn gorm.ConnPool) (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      Conn,
		DSN:                       dsn,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	return gormDB, err
}

func OpenDB() (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	return gormDB, err
}
