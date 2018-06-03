package core

import "github.com/gorilla/websocket"

type ClientManager struct {
	Clients []*Client
}

func (cm *ClientManager) registerConnection(conn *websocket.Conn) {
	cm.Clients = append(cm.Clients, &Client{Conn: conn})
}

func (cm *ClientManager) removeConnection(conn *websocket.Conn) {
	cm.removeUserByConn(conn)
}

// REALLY CRAPPY APPROACH, Replace with something efficient
func (cm *ClientManager) removeUserByConn(conn *websocket.Conn) {
	result := []*Client{}
	for _, usr := range cm.Clients {
		if usr.Conn != conn {
			result = append(result, usr)
		}
	}
	cm.Clients = result
}

// REALLY CRAPPY APPROACH, Replace with something efficient
func (cm *ClientManager) getUserByConn(conn *websocket.Conn) *Client {
	for _, usr := range cm.Clients {
		if usr.Conn == conn {
			return usr
		}
	}
	return nil
}
