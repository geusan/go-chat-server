package rest

import (
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

	e.GET("/users", handler.Fetch)
}

func (h *UserHandler) Fetch(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

// func getStatusCode(err error) int {
// 	if err == nil {
// 		return http.StatusOK
// 	}
// 	logrus.Error(err)
// 	switch err {
// 	case domain.ErrInternalServerError:
// 		return http.StatusInternalServerError
// 	case domain.ErrNotFound:
// 		return http.StatusNotFound
// 	case domain.ErrConflict:
// 		return http.StatusConflict
// 	default:
// 		return http.StatusInternalServerError
// 	}
// }
