package domain

import (
	"gorm.io/gorm"
)

type Chatroom struct {
	gorm.Model
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Limit   int    `json:"limit"`
	OwnerID uint   `json:"-"`
	Owner   User   `json:"-"`
	Users   []User `json:"users" gorm:"many2many:user_chatrooms;"`
}
