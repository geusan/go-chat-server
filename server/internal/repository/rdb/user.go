package rdb

import (
	domain "api-server/domain"
	"crypto/sha256"

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

func (m *UserRepository) Create(d *domain.User) *domain.User {
	res := m.Conn.Model(&domain.User{}).Create(&domain.User{Name: d.Name, Password: Salt(d.Password)})
	if res.Error != nil {
		panic(res.Error)
	}
	res.Scan(d)
	return d
}

func (m *UserRepository) Delete(id uint) error {
	res := m.Conn.Model(&domain.User{}).
		Unscoped().
		Where(&domain.User{ID: id}).
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
	return string(bs)
}
