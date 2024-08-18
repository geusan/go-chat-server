package main

import (
	"api-server/domain"
	"api-server/internal/repository/rdb"

	"github.com/sirupsen/logrus"
)

func main() {
	db, err := rdb.OpenDB()
	if err != nil {
		logrus.Error("error in DB", err)
		return
	}

	db.AutoMigrate(&domain.User{}, &domain.Chatroom{}, &domain.Chat{})
}
