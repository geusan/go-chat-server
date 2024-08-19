package rest_test

import (
	"api-server/domain"
	"api-server/internal/rest"
	"api-server/internal/rest/mocks"
	testing_utils "api-server/internal/utils"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"
)

func beforeTest(t *testing.T) (domain.User, *mocks.AuthService, *mocks.ChatService) {
	var mockAuth domain.User
	err := faker.FakeData(&mockAuth)

	assert.NoError(t, err)

	mockAuthService := new(mocks.AuthService)
	mockChatService := new(mocks.ChatService)

	return mockAuth, mockAuthService, mockChatService
}

func TestCreateChatroom(t *testing.T) {
	mockAuth, mockAuthService, mockChatService := beforeTest(t)

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

		err := handler.CreateChatroom(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockChatService.AssertExpectations(t)

	})
	// TODO: If adding validations,
	// t.Run("Error because blank Name", func(t *testing.T) {
	// })
}

func TestFetchChatrooms(t *testing.T) {
	_, mockAuthService, mockChatService := beforeTest(t)

	t.Run("Success", func(t *testing.T) {
		c, _, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.GET, "/v1/rooms", nil)
		expectedChatroom := []domain.Chatroom{
			{
				Name:  "chatoom 1",
				Limit: 5,
			},
			{
				Name:  "chatoom 2",
				Limit: 5,
			},
		}
		mockChatService.
			On("Fetch").
			Return(expectedChatroom).
			Once()

		handler := rest.ChatroomHandler{ChatService: mockChatService, UserService: mockAuthService}
		err := handler.Fetch(c)
		var actualChatrooms []domain.Chatroom
		json.NewDecoder(rec.Body).Decode(&actualChatrooms)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		reflect.DeepEqual(expectedChatroom, actualChatrooms)
		mockChatService.AssertExpectations(t)
	})
}

func TestRemoveChatroom(t *testing.T) {
	mockAuth, mockAuthService, mockChatService := beforeTest(t)

	t.Run("Success", func(t *testing.T) {
		roomId := uint(1)
		c, _, rec := testing_utils.NewTestHttp(t, context.TODO(), echo.DELETE, "/v1/rooms/"+strconv.Itoa(int(roomId)), nil)
		c.Set("auth", mockAuth)

		c.SetPath("/v1/rooms/:roomId")
		c.SetParamNames("roomId")
		c.SetParamValues(strconv.Itoa(int(roomId)))

		expectedChatroom := domain.Chatroom{
			ID:    roomId,
			Name:  "chatoom 1",
			Limit: 5,
		}

		mockChatService.
			On("FindById", roomId).
			Return(&expectedChatroom).
			Once()

		mockChatService.
			On("Delete", &expectedChatroom).
			Return(nil).
			Once()

		handler := rest.ChatroomHandler{ChatService: mockChatService, UserService: mockAuthService}
		err := handler.RemoveChatroom(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		mockChatService.AssertExpectations(t)
	})

	// t.Run("Failed by no permission", func(t *testing.T) {})

}
