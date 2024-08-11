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

func TestLogin(t *testing.T) {

}
