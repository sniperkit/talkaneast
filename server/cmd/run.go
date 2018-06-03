package main

import (
	_ "github.com/nneast/talkaneast/server/pkg/channels"
	"github.com/nneast/talkaneast/server/pkg/core"
	_ "github.com/nneast/talkaneast/server/pkg/messages"
	_ "github.com/nneast/talkaneast/server/pkg/users"
)

func main() {
	app := core.Application{}
	app.Run()
}
