package auth

import (
	"api-server/domain"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	FindOne(query *domain.User) (result *domain.User, err error)
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
	query := &domain.User{
		Name:     name,
		Password: password,
	}
	user, err := a.userRepo.FindOne(query)
	if err != nil {
		panic(err)
	}
	return user
}

func (a *AuthService) FindUserByName(name string) *domain.User {
	query := &domain.User{
		Name: name,
	}
	user, err := a.userRepo.FindOne(query)
	if err != nil {
		return nil
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
