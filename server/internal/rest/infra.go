package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type InfraHandler struct {
	Service ChatService
}

func NewInfraHandler(e *echo.Group, svc ChatService) {
	handler := &InfraHandler{
		Service: svc,
	}

	e.POST("/chat-servers", handler.AddServer)
}

func (h *InfraHandler) AddServer(c echo.Context) error {
	// TODO: Add authentication checker for only server
	var body struct {
		Url string `json:"url"`
	}
	ParseBody(c, &body)
	println("Add Server", body.Url)
	h.Service.AddServer(body.Url)
	return c.JSON(http.StatusCreated, "added")
}
