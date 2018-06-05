package messages

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	EVENT_MESSAGE = "Message"
)

type Message struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Content   string        `json:"content" bson:"content"`
	AuthorID  bson.ObjectId `json:"author_id" bson:"author_id"`
	ImageUrl  string        `json:"image_url" bson:"image_url"`
	CreatedOn time.Time     `json:"createdOn" structs:"createdOn" bson:"createdOn"`
}

/*func (server *core.Server) sendMessageToUser(receiver *users.User, sender *users.User, message string) {
	receiver.Conn.WriteJSON(CreateEvent(EVENT_MESSAGE, Message{Message: message, Username: sender.Username}))
}*/
