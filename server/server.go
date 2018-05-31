package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"time"
)

type Server struct {
	Users        []*User
	EventHandler *EventHandler
}

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (server *Server) handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	log.Println("Registered user " + strconv.Itoa(len(server.Users)))
	server.Users = append(server.Users, &User{Conn: c})

	defer c.Close()

	for {
		var response map[string]interface{}
		err := c.ReadJSON(&response)
		if err != nil {
			log.Println("Disconnected client")
			// Remove client
			server.removeUserByConn(c)
			break
		}

		server.EventHandler.handleEvent(server.getUserByConn(c), response)
	}
}

func (server *Server) startServer() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/ws", server.handleWS)
	log.Fatal(http.ListenAndServe(":2148", nil))
}

func main() {
	server := &Server{
		Users: make([]*User, 0),
	}

	server.EventHandler = &EventHandler{Server: server}

	server.startServer()
}

// REALLY CRAPPY APPROACH, Replace with something efficient
func (server *Server) removeUserByConn(conn *websocket.Conn) {
	result := []*User{}
	for _, usr := range server.Users {
		if usr.Conn != conn {
			result = append(result, usr)
		}
	}
	server.Users = result
}

// REALLY CRAPPY APPROACH, Replace with something efficient
func (server *Server) getUserByConn(conn *websocket.Conn) *User {
	for _, usr := range server.Users {
		if usr.Conn == conn {
			return usr
		}
	}
	return nil
}

func makeTimestamp() int64 {
	return time.Now().Unix() * int64(time.Millisecond)
}
