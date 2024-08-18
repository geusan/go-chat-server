package auth

import (
	"api-server/domain"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	FindOne(name string, password string) (result *domain.User, err error)
	FindById(id uint) (result *domain.User)
	Create(d *domain.User) *domain.User
	Delete(id uint) error
}

type AuthService struct {
	userRepo UserRepository
}

func NewAuthService(u UserRepository) *AuthService {
	return &AuthService{userRepo: u}
}

func (a *AuthService) FindUserByNameAndPassword(name string, password string) *domain.User {
	user, err := a.userRepo.FindOne(name, password)
	if err != nil {
		panic(err)
	}
	return user
}

func (a *AuthService) FindUserById(id uint) *domain.User {
	user := a.userRepo.FindById(id)
	return user
}

func (a *AuthService) Register(name string, password string) *domain.User {
	// TODO: add validating password with regex
	user := a.userRepo.Create(&domain.User{Name: name, Password: password})
	return user
}
