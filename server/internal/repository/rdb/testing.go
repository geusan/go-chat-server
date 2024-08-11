package rdb

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err, fmt.Errorf("an error '%s' was not expected when opening a stub database", err))

	if err != nil {
		defer db.Close()
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		DSN:                       "root:localhost@tcp(127.0.0.1:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	// gormDB, err := gorm.Open(sqlite.New(sqlite.Config{
	// 	Conn:       db,
	// 	DriverName: "sqlite3",
	// 	DSN:        "sqlite3.db",
	// }), &gorm.Config{})

	assert.NoError(t, err, fmt.Errorf("an error '%s' was not expected when opening gorm database", err))

	return gormDB, mock
}
