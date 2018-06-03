package core

import (
	"container/list"
	"log"
	"reflect"
)

var ControllersToRegister = list.New()

type EventHandler struct {
	App *Application
}

func (eh *EventHandler) registerHandlers() {
	for h := ControllersToRegister.Front(); h != nil; h = h.Next() {
		h.Value.(EventController).Register(eh.App)
		log.Printf("Registered event controller %s", reflect.TypeOf(h.Value).String())
	}
}

const waitTime int64 = 1000

func (eh *EventHandler) handleEvent(client *Client, event map[string]interface{}) {
	if _, ok := event["event"].(string); ok {
		for h := ControllersToRegister.Front(); h != nil; h = h.Next() {
			cRef := reflect.TypeOf(h.Value.(EventController))
			_, ok := cRef.MethodByName("Event" + event["event"].(string))

			if ok {
				args := []reflect.Value{reflect.ValueOf(client), reflect.ValueOf(event["data"].(map[string]interface{}))}
				cRef := reflect.ValueOf(h.Value.(EventController))
				method := cRef.MethodByName("Event" + event["event"].(string))
				method.Call(args)
			}
		}
	}
}
