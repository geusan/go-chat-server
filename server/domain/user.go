package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AddUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
