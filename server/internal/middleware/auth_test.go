package middleware_test

import (
	"api-server/domain"
	"api-server/internal/middleware"
	"api-server/internal/middleware/mocks"
	"api-server/internal/rest"
	testing_utils "api-server/internal/utils"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery --name AuthService
type AuthService interface {
	FindUserByNameAndPassword(name string, password string) *domain.User
	Register(name string, password string) *domain.User
	FindUserById(id uint) *domain.User
}

func TestUseAuthMiddleware(t *testing.T) {
	mockAuthService := new(mocks.AuthService)

	echojwtMiddlewareFunc, middlewareFunc := middleware.CreateAuthMiddlewareFunc(mockAuthService)

	h := echojwtMiddlewareFunc(middlewareFunc(echo.HandlerFunc(func(c echo.Context) error {
		user := c.Get("auth").(domain.User)
		return c.JSON(http.StatusOK, user)
	})))

	t.Run("200", func(t *testing.T) {
		mockUser := &domain.User{
			Id:   1,
			Name: "Luther",
		}
		token, err := mockUser.GenerateJWT()
		assert.NoError(t, err)
		mockAuthService.On("FindUserById", mockUser.Id).
			Return(mockUser).Once()

		c, req, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.GET, "/auth", nil)

		req.Header.Set("Authorization", "Bearer "+token)
		assert.NoError(t, h(c))

		actualUser := c.Get("auth").(domain.User)
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
		assert.Equal(t, mockUser.Id, actualUser.Id)
		assert.Equal(t, mockUser.Name, actualUser.Name)
		mockAuthService.AssertExpectations(t)
	})

	t.Run("401", func(t *testing.T) {
		c, req, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.GET, "/auth", nil)
		req.AddCookie(&http.Cookie{
			Name:  "_auth",
			Value: "Bearer XXX",
		})
		assert.NoError(t, h(c))
		var responseBody rest.ResponseError
		json.NewDecoder(rec.Body).Decode(&responseBody)
		assert.Contains(t, responseBody.Message, "token is malformed")
		assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
		mockAuthService.AssertNotCalled(t, "FindUserById")
	})

	t.Run("401 with no cookie", func(t *testing.T) {
		c, _, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.GET, "/auth", nil)
		assert.NoError(t, h(c))
		var responseBody rest.ResponseError
		json.NewDecoder(rec.Body).Decode(&responseBody)
		assert.Contains(t, responseBody.Message, "missing value in cookies")
		assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
		mockAuthService.AssertNotCalled(t, "FindUserById")
	})

	t.Run("400 malformed JWT", func(t *testing.T) {
		c, req, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.GET, "/auth", nil)
		req.Header.Set("Authorization", "BBearer XXXX")
		req.AddCookie(&http.Cookie{
			Name:  "_auth",
			Value: "Bearer XXX",
		})

		assert.NoError(t, h(c))

		var responseBody rest.ResponseError
		json.NewDecoder(rec.Body).Decode(&responseBody)
		assert.Contains(t, responseBody.Message, "token is malformed")
		assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
		mockAuthService.AssertNotCalled(t, "FindUserById")
	})
}
