package main

type Event struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func CreateEvent(event string, data interface{}) Event {
	return Event{
		Event: event,
		Data:  data,
	}
}
