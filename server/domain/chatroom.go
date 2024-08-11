package domain

import (
	"gorm.io/gorm"
)

type Chatroom struct {
	gorm.Model
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Limit   int    `json:"limit"`
	OwnerID uint   `json:"ownerId"`
	Owner   User   `json:"owner"`
	Users   []User `json:"users" gorm:"many2many:user_chatrooms;"`
}
