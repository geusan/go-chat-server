package rest

import (
	"api-server/domain"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//go:generate mockery --name ChatService
type ChatService interface {
	FindById(id uint) *domain.Chatroom
	Fetch() []domain.Chatroom
	Create(name string, user *domain.User) *domain.Chatroom
	Delete(chatroom *domain.Chatroom)
}

type ChatroomHandler struct {
	ChatService ChatService
	UserService AuthService
}

func NewChatroomHandler(e *echo.Group, svc ChatService, authService AuthService) {
	handler := &ChatroomHandler{
		ChatService: svc,
		UserService: authService,
	}

	e.GET("/rooms", handler.Fetch)
	e.POST("/rooms", handler.CreateChatroom)
	e.DELETE("/rooms/:roomId", handler.RemoveChatroom)
	e.GET("/rooms/:roomId/open", handler.OpenChat)
}

// Fetch chatrooms
//
// @Summary Fetch Chatrooms
// @Description Get list of chatrooms
// @Tags chat
// @Accept json
// @Produce json
// @Param	Authorization	header	string	true "Bearer XXX"
// @Success	200	{array}		[]domain.Chatroom
// @Success	400	{object}	ResponseError
// @Success	404	{object}	ResponseError
// @Success	500	{object}	ResponseError
// @Router       /rooms [get]
func (h *ChatroomHandler) Fetch(c echo.Context) error {
	chats := h.ChatService.Fetch()
	return c.JSON(http.StatusOK, chats)
}

type CreateChatroomDTO struct {
	Name string `json:"name" example:"new chatroom"`
}

// Create Chatroom
//
// @Summary Create chatroom
// @Description Create new chatroom
// @Tags chat
// @Accept json
// @Produce json
// @Param	Authorization	header	string	true "Bearer XXX"
// @Param	body	body	CreateChatroomDTO true "create chatroom dto"
// @Success	200	{array}		domain.Chatroom
// @Success	400	{object}	ResponseError
// @Success	404	{object}	ResponseError
// @Success	500	{object}	ResponseError
// @Router       /rooms [post]
func (h *ChatroomHandler) CreateChatroom(c echo.Context) error {
	user := c.Get("auth").(domain.User)
	var body CreateChatroomDTO
	body = ParseBody(c, body)
	chat := h.ChatService.Create(body.Name, &user)

	return c.JSON(http.StatusOK, *chat)
}

// Remove Chatroom
//
// @Summary Delete chatroom
// @Description Delete chatroom
// @Tags chat
// @Accept json
// @Produce json
// @Param	Authorization	header	string	true "Bearer XXX"
// @Param	roomId	path int true "delete chatroom id"
// @Success	200	{array}		domain.Chatroom
// @Success	400	{object}	ResponseError
// @Success	404	{object}	ResponseError
// @Success	500	{object}	ResponseError
// @Router       /rooms/{roomId} [delete]
func (h *ChatroomHandler) RemoveChatroom(c echo.Context) error {
	rawRoomId := c.Param("roomId")
	roomId, err := strconv.ParseInt(rawRoomId, 10, 64)
	if err != nil {
		panic(err)
	}

	chatroom := h.ChatService.FindById(uint(roomId))

	h.ChatService.Delete(chatroom)
	return c.JSON(http.StatusOK, "")
}

func (h *ChatroomHandler) OpenChat(c echo.Context) error {
	// TODO: make chat url
	// chatroom := c.Param("chatroom")
	// hub := h.ChatService.GetHub(chatroom)
	// openWebsocket(hub, c.Response().Writer, c.Request())
	return nil
}
