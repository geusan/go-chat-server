package rdb

import (
	"api-server/domain"

	"gorm.io/gorm"
)

type ChatroomRepository struct {
	Conn *gorm.DB
}

func NewChatroomRepository(conn *gorm.DB) *ChatroomRepository {
	return &ChatroomRepository{conn}
}

func (m *ChatroomRepository) FindById(id uint) (result *domain.Chatroom) {
	res := m.Conn.
		Model(&domain.Chatroom{}).
		Where(&domain.Chatroom{ID: id}).
		First(&result)

	if res.Error != nil {
		panic(res.Error)
	}
	return result
}

func (m *ChatroomRepository) Fetch() []domain.Chatroom {
	var results []domain.Chatroom
	result := m.Conn.
		Limit(10).
		Find(&results)

	if result.Error != nil {
		panic(result.Error)
	}

	return results
}

func (m *ChatroomRepository) Create(name string, owner *domain.User) *domain.Chatroom {
	chatroom := &domain.Chatroom{Name: name, Owner: *owner}
	res := m.Conn.Model(&domain.Chatroom{}).Create(&chatroom)
	if res.Error != nil {
		panic(res.Error)
	}
	res.Scan(chatroom)
	return chatroom
}

func (m *ChatroomRepository) Delete(id uint) error {
	res := m.Conn.Model(&domain.Chatroom{}).Delete(&domain.Chatroom{ID: id})
	if res.Error != nil {
		panic(res)
	}
	return nil
}
