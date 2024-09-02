package rest

import (
	"api-server/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	Service ChatService
}

func NewUserHandler(e *echo.Group, svc ChatService) {
	handler := &UserHandler{
		Service: svc,
	}

	e.GET("/me", handler.GetMe)
}

func (h *UserHandler) GetMe(c echo.Context) error {
	user := c.Get("auth").(domain.User)
	return c.JSON(http.StatusOK, &user)
}
