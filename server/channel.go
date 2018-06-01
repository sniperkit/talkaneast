package main

const (
	EVENT_CHANNELS_LIST = "ChannelsList"
)

type Channel struct {
	Name string `json:"name"`
}

func (server *Server) sendChannelsToUser(user *User) {
	user.Conn.WriteJSON(CreateEvent(EVENT_CHANNELS_LIST, server.Channels))
}
