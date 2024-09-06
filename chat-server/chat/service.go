package chat

//go:generate mockery --name ChatroomRepository
type ChatroomRepository interface{}

type ChatService struct {
	chatroomRepo ChatroomRepository
}

func NewChatService(c ChatroomRepository) *ChatService {
	return &ChatService{
		chatroomRepo: c,
	}
}

var hubMap = make(map[string]*Hub)

func (s *ChatService) GetOrCreateHub(chatroom string) *Hub {
	hub := createOrGetSocket(chatroom)
	return hub
}

func (s *ChatService) DeleteHub(chatroom string) *Hub {
	hub := hubMap[chatroom]
	if hub != nil {
		hub.Close()
	}
	return hub
}

func (s *ChatService) GetHub(chatroom string) *Hub {
	hub := hubMap[chatroom]
	return hub
}

func createOrGetSocket(chatroom string) *Hub {
	var hub *Hub
	hub = hubMap[chatroom]
	if hub == nil {
		hub = newHub()
	}

	go hub.run()
	hubMap[chatroom] = hub

	return hub
}
