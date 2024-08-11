package rdb_test

import (
	"api-server/domain"
	"api-server/internal/repository/rdb"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewUserRepository(db)
	user := &domain.User{
		Name:     "John",
		Password: "salt",
	}
	mock.ExpectBegin()
	mock.
		ExpectExec("INSERT INTO `users`").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	actualUser := mockRepo.Create(user)
	assert.Equal(t, user.Name, actualUser.Name)
}

func TestFindUser(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewUserRepository(db)
	name := "John"
	password := "salt"
	rows := sqlmock.
		NewRows([]string{"Name", "Password"}).
		AddRow(name, rdb.Salt(password))

	mock.
		ExpectQuery("^SELECT (.+) LIMIT ?").
		WithArgs(name, rdb.Salt(password), 1).
		WillReturnRows(rows)
	user, err := mockRepo.FindOne(name, password)
	assert.NoError(t, err)
	assert.Equal(t, name, user.Name)
}

func TestDeleteUser(t *testing.T) {
	db, mock := rdb.NewMockDB(t)
	mockRepo := rdb.NewUserRepository(db)

	mock.ExpectBegin()
	mock.
		ExpectExec("^DELETE FROM `users` WHERE (.+)$").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mockRepo.Delete(1)
}
