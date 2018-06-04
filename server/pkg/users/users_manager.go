package users

import (
	"github.com/NNeast/talkaneast/server/pkg/core"
	"github.com/gorilla/websocket"
)

type UsersManager struct {
	Users []*User
	App   *core.Application
}

func (um *UsersManager) RegisterConnection(conn *websocket.Conn) {
	//um.Users = append(um.Users, &User{Conn: conn})
}

func (um *UsersManager) RemoveConnection(conn *websocket.Conn) {
	//um.removeUserByConn(conn)
}
