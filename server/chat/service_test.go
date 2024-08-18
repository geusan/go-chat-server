package chat_test

import (
	"api-server/chat"
	"api-server/chat/mocks"
	"api-server/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	mockChatroomRepo := new(mocks.ChatroomRepository)
	mockUserRepo := new(mocks.UserRepository)
	var mockChatrooms []domain.Chatroom
	mockChatrooms = append(mockChatrooms, domain.Chatroom{
		Name:  "chatroom 1",
		Limit: 5,
	}, domain.Chatroom{
		Name:  "chatroom 2",
		Limit: 5,
	})

	t.Run("Success", func(t *testing.T) {
		mockChatroomRepo.
			On("Fetch").
			Return(mockChatrooms).
			Once()
		service := chat.NewChatService(mockUserRepo, mockChatroomRepo)
		actualChatrooms := service.Fetch()
		assert.Equal(t, len(mockChatrooms), len(actualChatrooms))
		mockChatroomRepo.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {
	mockChatroomRepo := new(mocks.ChatroomRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockChatroom := &domain.Chatroom{
		ID:    uint(1),
		Name:  "chatroom 1",
		Limit: 5,
	}

	t.Run("Success", func(t *testing.T) {
		mockChatroomRepo.
			On("FindById", mockChatroom.ID).
			Return(mockChatroom).
			Once()
		service := chat.NewChatService(mockUserRepo, mockChatroomRepo)
		actualChatroom := service.FindById(mockChatroom.ID)
		assert.Equal(t, mockChatroom.ID, actualChatroom.ID)
		mockChatroomRepo.AssertExpectations(t)
	})
}

func TestCreateChatroom(t *testing.T) {
	mockChatroomRepo := new(mocks.ChatroomRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockUser := &domain.User{
		Id:       1,
		Name:     "John",
		Password: "Salt",
	}
	mockChatroom := &domain.Chatroom{
		Name:  "chatroom 1",
		Limit: 5,
		Owner: *mockUser,
	}

	t.Run("Success", func(t *testing.T) {
		mockChatroomRepo.
			On("Create", mockChatroom.Name, mockUser).
			Return(mockChatroom).
			Once()
		service := chat.NewChatService(mockUserRepo, mockChatroomRepo)
		actualChatroom := service.Create(mockChatroom.Name, mockUser)
		assert.Equal(t, mockChatroom.Owner.Id, actualChatroom.Owner.Id)
		assert.Equal(t, mockChatroom.Name, actualChatroom.Name)
		mockChatroomRepo.AssertExpectations(t)
	})
}

func TestDeleteChatroom(t *testing.T) {
	mockChatroomRepo := new(mocks.ChatroomRepository)
	mockUserRepo := new(mocks.UserRepository)
	mockChatroom := &domain.Chatroom{
		ID:    1,
		Name:  "chatroom 1",
		Limit: 5,
	}

	t.Run("Success", func(t *testing.T) {
		mockChatroomRepo.
			On("Delete", mockChatroom.ID).
			Return(nil).
			Once()
		service := chat.NewChatService(mockUserRepo, mockChatroomRepo)
		service.Delete(mockChatroom)
		mockChatroomRepo.AssertExpectations(t)
	})
}
