package chat

import (
	"api-server/domain"
	"fmt"
)

//go:generate mockery --name UserRepository
type UserRepository interface {
	FindOne(query *domain.User) (result *domain.User, err error)
	Create(d *domain.User) *domain.User
	Delete(id uint) error
}

//go:generate mockery --name ChatroomRepository
type ChatroomRepository interface {
	FindById(id uint) *domain.Chatroom
	Fetch() []domain.Chatroom
	Create(name string, owner *domain.User) *domain.Chatroom
	Delete(id uint) error
}

//go:generate mockery --name ChatroomHashRepository
type ChatroomHashRepository interface {
	AddServer(url string)
	GetServer(key string) string
}

type ChatService struct {
	userRepo     UserRepository
	chatroomRepo ChatroomRepository
	ccr          ChatroomHashRepository
}

func NewChatService(u UserRepository, c ChatroomRepository, r ChatroomHashRepository) *ChatService {
	return &ChatService{
		userRepo:     u,
		chatroomRepo: c,
		ccr:          r,
	}
}

func (s *ChatService) Fetch() []domain.Chatroom {
	chatrooms := s.chatroomRepo.Fetch()
	return chatrooms
}

func (s *ChatService) FindById(id uint) *domain.Chatroom {
	chatroom := s.chatroomRepo.FindById(id)
	return chatroom
}

func (s *ChatService) Create(name string, user *domain.User) *domain.Chatroom {
	chatroom := s.chatroomRepo.Create(name, user)
	return chatroom
}

func (s *ChatService) Delete(chatroom *domain.Chatroom) {
	s.chatroomRepo.Delete(chatroom.ID)
}

func (s *ChatService) Open(chatroom *domain.Chatroom, user *domain.User) string {
	// Save redis for auth
	member := s.ccr.GetServer(fmt.Sprint(chatroom.Id))
	return member
}

func (s *ChatService) AddServer(url string) {
	s.ccr.AddServer(url)
}
