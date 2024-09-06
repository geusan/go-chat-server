package rest

import (
	"chat-server/chat"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name ChatService
type ChatService interface {
	GetHub(chatroom string) *chat.Hub
	GetOrCreateHub(chatroom string) *chat.Hub
	DeleteHub(chatroom string) *chat.Hub
}

type ChatroomHandler struct {
	ChatService ChatService
}

func NewChatroomHandler(e *echo.Group, svc ChatService) {
	handler := &ChatroomHandler{
		ChatService: svc,
	}

	e.GET("/rooms/:roomId/open", handler.OpenChat)
	e.GET("/rooms/:roomId/close", handler.CloseChat)
}

func (h *ChatroomHandler) OpenChat(c echo.Context) error {
	chatroom := c.Param("roomId")
	hub := h.ChatService.GetOrCreateHub(chatroom)
	if hub.GetCount() < 5 {
		openWebsocket(hub, c.Response(), c.Request())
		return nil
	}
	return c.HTML(http.StatusPaymentRequired, "")
}

func (h *ChatroomHandler) CloseChat(c echo.Context) error {
	chatroom := c.Param("roomId")
	secret := c.Request().Header.Get("HTTP_SECRET")
	if secret != os.Getenv("SECRET") {
		return c.HTML(http.StatusBadRequest, "")
	}
	hub := h.ChatService.GetHub(chatroom)
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
