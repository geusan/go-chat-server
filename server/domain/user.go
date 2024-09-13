package domain

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique"`
	Password string `json:"-"`
}

type AddUser struct {
	Name     string `json:"name" example:"Luther"`
	Password string `json:"password" example:"umbrella"`
}

type ResponseUser struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

func (u *User) GenerateJWT() (string, error) {
	claim := new(JwtCustomClaims)
	claim.UserId = u.Id
	claim.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(7) * time.Hour))

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	s, err := t.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return s, err
}
