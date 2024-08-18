package rdb

import (
	domain "api-server/domain"
	"crypto/sha256"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{conn}
}

func (m *UserRepository) FindOne(name string, password string) (result *domain.User, err error) {
	res := m.Conn.
		Model(&domain.User{}).
		Where(&domain.User{Name: name, Password: Salt(password)}).
		First(&result)
	if res.Error != nil {
		panic(res.Error)
	}
	return result, nil
}

func (m *UserRepository) FindById(id uint) (result *domain.User) {
	res := m.Conn.
		Model(&domain.User{}).
		Where(&domain.User{Id: id}).
		First(&result)
	if res.Error != nil {
		panic(res.Error)
	}
	return result
}

func (m *UserRepository) Create(d *domain.User) *domain.User {
	d.Password = Salt(d.Password)
	res := m.Conn.Model(&domain.User{}).Create(d)
	if res.Error != nil {
		panic(res.Error)
	}
	res.Scan(d)
	return d
}

func (m *UserRepository) Delete(id uint) error {
	res := m.Conn.Model(&domain.User{}).
		Unscoped().
		Where(&domain.User{Id: id}).
		Delete(&domain.User{})
	if res.Error != nil {
		panic(res.Error)
	}
	return nil
}

func Salt(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
