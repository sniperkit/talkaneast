package core

import (
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	CurrentChannelID bson.ObjectId `json:"currentChannelID" structs:"currentChannelID" bson:"currentChannelID,omitempty"`
	UserID           bson.ObjectId
	Conn             *websocket.Conn
}

func (c *Client) SendEvent(event *Event) {
	c.Conn.WriteJSON(event)
}
