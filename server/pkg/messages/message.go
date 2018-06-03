package messages

import "gopkg.in/mgo.v2/bson"

const (
	EVENT_MESSAGE = "Message"
)

type Message struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Content  string        `json:"content" bson:"content"`
	AuthorID string        `json:"author_id" bson:"author_id"`
}

/*func (server *core.Server) sendMessageToUser(receiver *users.User, sender *users.User, message string) {
	receiver.Conn.WriteJSON(CreateEvent(EVENT_MESSAGE, Message{Message: message, Username: sender.Username}))
}*/
