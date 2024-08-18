package rdb_test

import (
	"api-server/domain"
	"api-server/internal/repository/rdb"
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFetchChatroom(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewChatroomRepository(db)

	rows := sqlmock.
		NewRows([]string{"Name", "Limit"}).
		AddRows([][]driver.Value{{"NewChatRoom1", 5}, {"NewChatRoom2", 5}}...)
	mock.
		ExpectQuery(`^SELECT (.+) LIMIT ?`).
		WithArgs(10).
		WillReturnRows(rows)
	actualResults := mockRepo.Fetch()

	assert.Equal(t, 2, len(actualResults))
}

func TestFindMyId(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewChatroomRepository(db)
	testId := uint(1)

	rows := sqlmock.
		NewRows([]string{"ID", "Name", "Limit"}).
		AddRow(testId, "new chatroom", 5)

	mock.
		ExpectQuery("^SELECT (.+) WHERE (.+) LIMIT ?").
		WithArgs(testId, 1).
		WillReturnRows(rows)
	chatroom := mockRepo.FindById(testId)
	assert.Equal(t, testId, chatroom.ID)
}

func TestCreateChatroom(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewChatroomRepository(db)
	chatroomName := "new chatroom"
	user := &domain.User{
		Id:       1,
		Name:     "John",
		Password: "Salt",
	}
	mock.ExpectBegin()
	// User가 없으면 User부터 생성함
	mock.
		ExpectExec("^INSERT INTO `users`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO `chatrooms`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	chatroom := mockRepo.Create(chatroomName, user)
	assert.Equal(t, chatroomName, chatroom.Name)
	assert.Equal(t, user.Name, chatroom.Owner.Name)
}

func TestDeleteChatroom(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewChatroomRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM `chatrooms` WHERE (.+)$").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mockRepo.Delete(1)
}
