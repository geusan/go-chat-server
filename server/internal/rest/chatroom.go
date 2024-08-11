package rest

import (
	"api-server/chat"
	"api-server/domain"
	"api-server/internal/middleware"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ChatService interface {
	Fetch(ctx context.Context) []domain.Chatroom
	GetHub(chatroom string) *chat.Hub
	Create(name string, user *domain.User) *domain.Chatroom
	Delete(chatroom *domain.Chatroom)
}

type ChatroomHandler struct {
	Service ChatService
}

func NewChatroomHandler(e *echo.Group, svc ChatService) {
	handler := &ChatroomHandler{
		Service: svc,
	}
	router := e.Group("/rooms")
	middleware.UseAuthMiddleware(router)
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("Hey~~~")
			next(c)
			return nil
		}
	})
	router.GET("/rooms", handler.Fetch)
	router.POST("/rooms", handler.CreateChatroom)
	router.GET("/rooms/:chatroom/open", handler.OpenChat)
	router.GET("/rooms/:chatroom/close", handler.CloseChat)
}

func (h *ChatroomHandler) Fetch(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func (h *ChatroomHandler) CreateChatroom(c echo.Context) error {
	// user := c.Get("user").(*jwt.Claims)
	var body struct {
		name string
	}
	ParseBody(c, body)

	// h.Service.Create()
	return nil
}

func (h *ChatroomHandler) RemoveChatroom(c echo.Context) error {
	return nil
}

func (h *ChatroomHandler) OpenChat(c echo.Context) error {
	chatroom := c.Param("chatroom")
	hub := h.Service.GetHub(chatroom)
	openWebsocket(hub, c.Response().Writer, c.Request())
	return nil
}

func (h *ChatroomHandler) CloseChat(c echo.Context) error {
	chatroom := c.Param("chatroom")
	hub := h.Service.GetHub(chatroom)
	hub.Close()
	return nil
}

func openWebsocket(hub *chat.Hub, w http.ResponseWriter, r *http.Request) {

	conn, err := chat.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := chat.NewChatClient(hub, conn, make(chan []byte, 256))
	client.Hub.AddClient(client)

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
