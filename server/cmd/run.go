package main

import (
	_ "github.com/NNeast/talkaneast/server/pkg/channels"
	"github.com/NNeast/talkaneast/server/pkg/core"
	_ "github.com/NNeast/talkaneast/server/pkg/messages"
	_ "github.com/NNeast/talkaneast/server/pkg/users"
)

func main() {
	app := core.Application{}
	app.Run()
}
