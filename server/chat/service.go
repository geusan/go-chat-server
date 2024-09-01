package chat

import (
	"api-server/domain"
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
