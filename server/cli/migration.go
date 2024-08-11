package main

import (
	"api-server/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite3.db"), &gorm.Config{})
	if err != nil {
		logrus.Error("error in DB", err)
		return
	}

	db.AutoMigrate(&domain.User{}, &domain.Chatroom{}, &domain.Chat{})
}
