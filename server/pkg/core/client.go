package core

import "github.com/gorilla/websocket"

type Client struct {
	UserID string
	Conn   *websocket.Conn
}

func (c *Client) SendEvent(event *Event) {
	c.Conn.WriteJSON(event)
}
