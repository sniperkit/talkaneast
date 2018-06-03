package core

import (
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

type Client struct {
	UserID bson.ObjectId
	Conn   *websocket.Conn
}

func (c *Client) SendEvent(event *Event) {
	c.Conn.WriteJSON(event)
}
