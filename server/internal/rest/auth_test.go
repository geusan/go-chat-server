package rest_test

import (
	"api-server/domain"
	"api-server/internal/rest"
	"api-server/internal/rest/mocks"
	testing_utils "api-server/internal/utils"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	faker "github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	// 테스트 준비
	var mockAuth domain.User
	err := faker.FakeData(&mockAuth)
	assert.NoError(t, err)
	mockService := new(mocks.AuthService)
	// 테스트 로직 시작
	username := "john"
	password := "password"
	// service에서 함수 호출이 잘 되었고 리턴 값이 예상대로인지
	mockService.On("Register", username, password).Return(&domain.User{Name: username, Password: password})

	body := &domain.AddUser{Name: username, Password: password}
	c, req, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.POST, "/register", body)
	req.Header.Set("Content-Type", "application/json")

	handler := rest.AuthHandler{Service: mockService}
	err = handler.Register(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	// 테스트 준비
	var mockAuth domain.User
	err := faker.FakeData(&mockAuth)
	assert.NoError(t, err)
	mockService := new(mocks.AuthService)
	// 테스트 로직 시작
	username := "john"
	password := "password"
	// service에서 함수 호출이 잘 되었고 리턴 값이 예상대로인지
	mockService.On("Login", username, password).Return(&domain.User{Name: username, Password: password})

	body := &domain.AddUser{Name: username, Password: password}
	buf, err := json.Marshal(body)
	assert.NoError(t, err)
	c, _, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.POST, "/register", bytes.NewBuffer(buf))

	handler := rest.AuthHandler{Service: mockService}
	err = handler.Login(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}
