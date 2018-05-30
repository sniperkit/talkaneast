package main

const (
	EVENT_NOTIFICATION = "Notification"
)

type Notification struct {
	NotificationMessage string `json:"message"`
}

func (server *Server) sendNotification(user *User, notification string) {
	user.Conn.WriteJSON(CreateEvent(EVENT_NOTIFICATION, Notification{NotificationMessage: notification}))
}
