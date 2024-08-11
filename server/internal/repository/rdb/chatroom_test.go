package rdb_test

import (
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
		ExpectQuery(`^SELECT (.+) OFFSET ?$`).
		WithArgs(10, 10).
		WillReturnRows(rows)
	actualResults := mockRepo.Fetch()

	assert.Equal(t, 2, len(actualResults))
}
