package domain

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Id       uint   `json:"id"`
	Content  string `json:"content"`
	WriterId uint   `json:"writerId"`
}
