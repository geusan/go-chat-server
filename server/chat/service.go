package chat

import (
	"context"

	"api-server/domain"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	FindOne(name string, password string) (result *domain.User, err error)
	Create(d *domain.User) *domain.User
}

//go:generate mockery --name ChatroomRepository
type ChatroomRepository interface {
	GetChatroom(id uint) *domain.Chatroom
	Fetch() []domain.Chatroom
	Create(name string, owner *domain.User) *domain.Chatroom
	Delete(id uint) error
}

type ChatService struct {
	userRepo     UserRepository
	chatroomRepo ChatroomRepository
}

func NewChatService(u UserRepository, c ChatroomRepository) *ChatService {
	return &ChatService{
		userRepo:     u,
		chatroomRepo: c,
	}
}

func (s *ChatService) Fetch(ctx context.Context) []domain.Chatroom {
	chatrooms := s.chatroomRepo.Fetch()
	return chatrooms
}

var hubMap = make(map[string]*Hub)

func (s *ChatService) GetHub(chatroom string) *Hub {
	hub := createOrGetSocket(chatroom)
	return hub
}

func (s *ChatService) Create(name string, user *domain.User) *domain.Chatroom {
	chatroom := s.chatroomRepo.Create(name, user)
	return chatroom
}

func (s *ChatService) Delete(chatroom *domain.Chatroom) {
	hub := hubMap[chatroom.Name]
	if hub != nil {
		hub.Close()
	}
	s.chatroomRepo.Delete(chatroom.ID)
}

func createOrGetSocket(chatroom string) *Hub {
	var hub *Hub
	hub = hubMap[chatroom]
	if hub == nil {
		hub = newHub()
	}

	go hub.run()
	hubMap[chatroom] = hub

	return hub
}
