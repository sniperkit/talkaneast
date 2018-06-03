package users

import (
	"github.com/gorilla/websocket"
	"github.com/nneast/talkaneast/server/pkg/core"
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

/*
// REALLY CRAPPY APPROACH, Replace with something efficient
func (um *UsersManager) removeUserByConn(conn *websocket.Conn) {
	result := []*User{}
	for _, usr := range um.Users {
		if usr.Conn != conn {
			result = append(result, usr)
		}
	}
	server.Users = result
}

// REALLY CRAPPY APPROACH, Replace with something efficient
func (um *UsersManager) getUserByConn(conn *websocket.Conn) *User {
	for _, usr := range um.Users {
		if usr.Conn == conn {
			return usr
		}
	}
	return nil
}

func (um *UsersManager) makeTimestamp() int64 {
	return time.Now().Unix() * int64(time.Millisecond)
}

func (um *UsersManager) containsUsers(s []*User, x string) bool {
	for _, a := range s {
		if a.Username == x {
			return true
		}
	}
	return false
}*/
