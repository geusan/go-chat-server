package consistent_hash

import (
	"api-server/domain"

	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash/v2"
)

type Hasher struct{}

func (h Hasher) Sum64(data []byte) uint64 {
	// you should use a proper hash function for uniformity.
	return xxhash.Sum64(data)
}

func CreateRing() *consistent.Consistent {
	config := consistent.Config{
		PartitionCount:    3,
		ReplicationFactor: 20,
		Load:              1.25,
		Hasher:            Hasher{},
	}
	c := consistent.New(nil, config)
	return c
}

type ChatroomHashRepository struct {
	ConsistentRing *consistent.Consistent
}

type ChatUser struct {
	User     *domain.User     `json:"user"`
	Chatroom *domain.Chatroom `json:"chatroom"`
}

func NewChatroomHashRepository(conn *consistent.Consistent) *ChatroomHashRepository {
	return &ChatroomHashRepository{
		ConsistentRing: conn,
	}
}

type Member struct {
	consistent.Member
	Url string
}

func (m Member) String() string {
	return m.Url
}

func (r *ChatroomHashRepository) AddServer(url string) {
	member := Member{Url: url}
	r.ConsistentRing.Add(member)
}

func (r *ChatroomHashRepository) GetServer(key string) string {
	_key := []byte(key)
	member := r.ConsistentRing.LocateKey(_key)
	return member.String()
}
