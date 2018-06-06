package core

import "github.com/fatih/structs"

type SocketError struct {
	Message string
	Code    int
}

func CreateError(message string, code int) map[string]interface{} {
	return structs.Map(SocketError{Message: message, Code: code})
}
