package chat_test

import (
	"go-swe/chapter04/chat"
	"reflect"
	"testing"
)

type entry struct {
	user    string
	message string
}

type spyPublisher struct {
	published []entry
}

func (p *spyPublisher) Publish(user, message string) error {
	p.published = append(p.published, entry{user: user, message: message})
	return nil
}

func TestChatRoomBroadcast(t *testing.T) {
	pub := new(spyPublisher)
	room := chat.NewRoom(pub)
	room.AddUser("bob")
	room.AddUser("alice")
	_ = room.Broadcast("hi")

	// check published entries
	expected := []entry{
		{user: "bob", message: "hi"},
		{user: "alice", message: "hi"},
	}

	if got := pub.published; !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected the following messages:\n%+v\ngot:\n%#+v", expected, got)
	}

}
