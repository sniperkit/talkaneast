package core

import "github.com/fatih/structs"

type SocketError struct {
	Message string
	Code    int
}

func CreateSocketError(message string, code int) map[string]interface{} {
	return structs.Map(SocketError{Message: message, Code: code})
}
