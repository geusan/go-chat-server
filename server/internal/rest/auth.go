package rest

import (
	"api-server/domain"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name AuthService
type AuthService interface {
	FindUserByNameAndPassword(name string, password string) *domain.User
	Register(name string, password string) *domain.User
	FindUserById(id uint) *domain.User
	FindUserByName(name string) *domain.User
	CheckPassword(user *domain.User, password string) bool
}

type AuthHandler struct {
	Service AuthService
}

func NewAuthHandler(e *echo.Group, svc AuthService) {
	handler := &AuthHandler{Service: svc}

	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
}

// Login user
//
// @Summary      Login
// @Description  login with id and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body    body     domain.AddUser  true  "name"
// @Success      200  {object}   domain.ResponseUser
// @Failure      400  {object}  ResponseError
// @Failure      404  {object}  ResponseError
// @Failure      500  {object}  ResponseError
// @Header       200  {string}  Authorization  "Bearer XXX"
// @Router       /login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	var body domain.AddUser
	if err := c.Bind(&body); err != nil {
		return err
	}
	user := h.Service.FindUserByName(body.Name)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, ResponseError{
			Message: "없는 유저입니다.",
		})
	}
	authorized := h.Service.CheckPassword(user, body.Password)
	if !authorized {
		return c.JSON(http.StatusUnauthorized, ResponseError{
			Message: "비밀번호가 틀렸습니다.",
		})
	}
	user.Password = ""
	token, err := user.GenerateJWT()
	if err != nil {
		println("error is occured : ", err)
	}
	// Set cookie for Auth
	cookie := new(http.Cookie)
	cookie.Name = "_auth"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Domain = "localhost"
	cookie.Path = "/"
	cookie.HttpOnly = true
	// TODO: In production mode set SameSite
	cookie.SameSite = http.SameSiteNoneMode
	cookie.Secure = true

	c.SetCookie(cookie)
	responseBody := &domain.ResponseUser{
		Id:    user.Id,
		Name:  user.Name,
		Token: cookie.Value,
	}
	return c.JSON(http.StatusOK, responseBody)
}

// Register User
//
//	@Summary      Register
//	@Description  create user with name and password
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
	exsist := h.Service.FindUserByName(body.Name)
	if exsist != nil {
		return c.JSON(http.StatusUnauthorized, ResponseError{
			Message: "이미 있습니다.",
		})
	}
	user := h.Service.Register(body.Name, body.Password)
	user.Password = ""
	return c.JSON(http.StatusOK, user)
}
