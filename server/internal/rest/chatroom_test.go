package rest_test

import (
	"api-server/domain"
	"api-server/internal/rest"
	"api-server/internal/rest/mocks"
	testing_utils "api-server/internal/utils"
	"context"
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"
)

func TestCreateChatroom(t *testing.T) {
	var mockAuth domain.User
	err := faker.FakeData(&mockAuth)

	assert.NoError(t, err)

	mockAuthService := new(mocks.AuthService)
	mockChatService := new(mocks.ChatService)

	t.Run("Success", func(t *testing.T) {
		var body rest.CreateChatroomDTO
		body.Name = "new_chat"

		c, req, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.POST, "/v1/rooms", body)
		req.Header.Set("Content-Type", "application/json")
		// Set User
		c.Set("auth", mockAuth)

		mockChatService.
			On("Create", body.Name, &mockAuth).
			Return(&domain.Chatroom{Name: body.Name, Owner: mockAuth}).
			Once()

		handler := rest.ChatroomHandler{ChatService: mockChatService, UserService: mockAuthService}

		err = handler.CreateChatroom(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockChatService.AssertExpectations(t)

	})
	// TODO: If adding validations,
	// t.Run("Error because blank Name", func(t *testing.T) {
	// })
}
