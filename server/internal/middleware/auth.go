package middleware

import (
	"api-server/auth"
	"api-server/domain"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func UseAuthMiddleware(e *echo.Group, authService *auth.AuthService) {
	config := echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "header:Authorization:Bearer ,cookie:_auth",
		ErrorHandler: func(c echo.Context, err error) error {
			println("ErrorHandler", err.Error())
			return err
		},
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
