package domain

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	ID       uint   `json:"id"`
	Content  string `json:"content"`
	Writer   User   `json:"writer"`
	WriterID uint   `json:"writerId"`
}
