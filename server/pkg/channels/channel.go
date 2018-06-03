package channels

import (
	"github.com/nneast/talkaneast/server/pkg/messages"
	"gopkg.in/mgo.v2/bson"
)

const (
	EVENT_CHANNELS_LIST = "ChannelsList"
)

type Channel struct {
	ID       bson.ObjectId       `json:"id" bson:"_id,omitempty"`
	Messages []*messages.Message `json:"messages" bson:"messages,omitempty"`
	Name     string              `json:"name" bson:"name"`
}
