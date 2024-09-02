package rdb

import (
	"gorm.io/gorm"
)

type ChatroomRepository struct {
	Conn *gorm.DB
}

func NewChatRepository(conn *gorm.DB) *ChatroomRepository {
	return &ChatroomRepository{conn}
}
