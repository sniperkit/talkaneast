package main

import (
	"log"
	"reflect"
)

type EventHandler struct {
	Server *Server
}

const waitTime int64 = 1000

func (eh *EventHandler) handleEvent(user *User, event map[string]interface{}) {
	if _, ok := event["event"].(string); ok {

		ehRef := reflect.ValueOf(eh)
		method := ehRef.MethodByName("Event" + event["event"].(string))
		args := []reflect.Value{reflect.ValueOf(user), reflect.ValueOf(event["data"].(map[string]interface{}))}
		method.Call(args)
	}
}

func (eh *EventHandler) EventMessage(user *User, data map[string]interface{}) {

	if user.Username == "" {
		eh.Server.sendNotification(user, "Set username first")
	} else if makeTimestamp() < (user.lastMessageMilis + waitTime) {
		eh.Server.sendNotification(user, "You have to wait before sending new message")
	} else {
		for _, usr := range eh.Server.Users {
			eh.Server.sendMessageToUser(usr, user, data["message"].(string))
			user.lastMessageMilis = makeTimestamp()
		}
	}
}

func (eh *EventHandler) EventSetNick(user *User, data map[string]interface{}) {
	if contains(eh.Server.Users, data["nickname"].(string)) {
		eh.Server.sendNotification(user, "this user already exist")
	}
	log.Printf("event set nick")
	user.Username = data["nickname"].(string)
}
