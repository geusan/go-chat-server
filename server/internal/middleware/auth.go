package middleware

import (
	"api-server/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

//go:generate mockery --name AuthService
type AuthService interface {
	FindUserByNameAndPassword(name string, password string) *domain.User
	Register(name string, password string) *domain.User
	FindUserById(id uint) *domain.User
}

func UseAuthMiddleware(e *echo.Group, authService AuthService) {
	config := echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization:Bearer ,cookie:_auth",
	}
	e.Use(echojwt.WithConfig(config), func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := &domain.JwtCustomClaims{}
			jwt.ParseWithClaims(token.Raw, claims, config.KeyFunc)
			user := authService.FindUserById(claims.UserId)
			// add user to context
			c.Set("auth", *user)
			return next(c)
		}
	})
}
