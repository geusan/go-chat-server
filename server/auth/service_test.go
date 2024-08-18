package auth_test

import (
	"api-server/auth"
	"api-server/auth/mocks"
	"api-server/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	name := "John"
	password := "salt"
	mockUser := &domain.User{
		Name:     name,
		Password: password,
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.
			On("Create", mockUser).
			Return(mockUser).
			Once()
		service := auth.NewAuthService(mockUserRepo)
		actualUser := service.Register(name, password)
		assert.Equal(t, name, actualUser.Name)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestFindUserById(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	name := "John"
	password := "salt"
	mockUser := &domain.User{
		Id:       uint(1),
		Name:     name,
		Password: password,
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.
			On("FindById", mockUser.Id).
			Return(mockUser).
			Once()
		service := auth.NewAuthService(mockUserRepo)
		acturalUser := service.FindUserById(mockUser.Id)
		assert.Equal(t, mockUser.Id, acturalUser.Id)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestFindUserByNameAndPassword(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	name := "John"
	password := "salt"
	mockUser := &domain.User{
		Id:       uint(1),
		Name:     name,
		Password: password,
	}

	t.Run("Success", func(t *testing.T) {
		mockUserRepo.
			On("FindOne", mockUser.Name, mockUser.Password).
			Return(mockUser, nil).
			Once()
		service := auth.NewAuthService(mockUserRepo)
		acturalUser := service.FindUserByNameAndPassword(mockUser.Name, mockUser.Password)
		assert.Equal(t, mockUser.Name, acturalUser.Name)
		assert.Equal(t, mockUser.Password, acturalUser.Password)
		mockUserRepo.AssertExpectations(t)
	})
}
