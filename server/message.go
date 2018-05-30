package main

const (
	EVENT_MESSAGE = "Message"
)

type Message struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

func (server *Server) sendMessageToUser(receiver *User, sender *User, message string) {
	receiver.Conn.WriteJSON(CreateEvent(EVENT_MESSAGE, Message{Message: message, Username: sender.Username}))
}
