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

func (m *ChatroomRepository) GetChatroom(id uint) *domain.Chatroom {
	result := m.Conn.Model(&domain.Chatroom{}).Find(&domain.Chatroom{ID: id})
	if result.Error != nil {
		panic(result.Error)
	}
	var chatroom domain.Chatroom
	row := result.Row()
	err := row.Scan(&chatroom)
	if err != nil {
		panic(err)
	}
	return &chatroom
}

func (m *ChatroomRepository) Fetch() []domain.Chatroom {
	result := m.Conn.Model(&domain.Chatroom{}).
		Limit(10).
		Offset(10).
		Find(&domain.Chatroom{})

	if result.Error != nil {
		panic(result.Error)
	}
	rows, err := result.Rows()
	if err != nil {
		panic(err)
	}
	results := make([]domain.Chatroom, 0)
	for rows.Next() {
		t := domain.Chatroom{}
		err = rows.Scan(&t)
		if err != nil {
			panic(err)
		}
		results = append(results, t)
	}
	return results
}

func (m *ChatroomRepository) Create(name string, owner *domain.User) *domain.Chatroom {
	chatroom := domain.Chatroom{Name: name, Owner: *owner}
	res := m.Conn.Model(&domain.Chatroom{}).Create(&chatroom)
	if res.Error != nil {
		panic(res.Error)
	}
	res.Scan(chatroom)
	return &chatroom
}

func (m *ChatroomRepository) Delete(id uint) error {
	res := m.Conn.Model(&domain.Chatroom{}).Delete(&domain.Chatroom{ID: id})
	if res.Error != nil {
		panic(res)
	}
	return nil
}
