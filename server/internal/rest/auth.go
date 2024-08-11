package rest

import (
	"api-server/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name AuthService
type AuthService interface {
	Login(name string, password string) *domain.User
	Register(name string, password string) *domain.User
}

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(e *echo.Group, svc AuthService) {
	handler := &AuthHandler{Service: svc}

	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
}

// ListAccounts lists all existing accounts
//
//	@Summary      Login
//	@Description  get accounts
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param        body    body     domain.AddUser  true  "name"
//	@Success      200  {object}   domain.User
//	@Failure      400  {object}  ResponseError
//	@Failure      404  {object}  ResponseError
//	@Failure      500  {object}  ResponseError
//	@Router       /login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var body struct {
		name     string
		password string
	}
	if err := c.Bind(&body); err != nil {
		return err
	}
	user := h.Service.Login(body.name, body.password)
	user.Password = ""
	return c.JSON(http.StatusOK, user)
}

// ListAccounts lists all existing accounts
//
//	@Summary      Register
//	@Description  get accounts
//	@Tags         auth
//	@Accept       json
//	@Produce      json
//	@Param        body    body     domain.AddUser  true  "name"
//	@Success      200  {object}   domain.User
//	@Failure      400  {object}  ResponseError
//	@Failure      404  {object}  ResponseError
//	@Failure      500  {object}  ResponseError
//	@Router       /register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	var body domain.AddUser
	if err := c.Bind(&body); err != nil {
		return err
	}

	println(body.Name, body.Password)
	user := h.Service.Register(body.Name, body.Password)
	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
