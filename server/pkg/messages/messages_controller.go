package messages

import (
	"github.com/NNeast/talkaneast/server/pkg/core"
)

type MessageController struct {
	app *core.Application
}

func (mc *MessageController) Register(app *core.Application) {
	mc.app = app
}

/*
func (mc *MessageController) EventMessage(user *services.User, data map[string]interface{}) {
	if user.Username == "" {
		//mc.app.sendNotification(user, "Set username first")
	} else if makeTimestamp() < (user.lastMessageMilis + waitTime) {
		//mc.app.sendNotification(user, "You have to wait before sending new message")
	} else {
		for _, usr := range mc.app.um.Users {
			if usr.CurrentChannel.Name == user.CurrentChannel.Name {
				mc.app.sendMessageToUser(usr, user, data["message"].(string))
				user.lastMessageMilis = makeTimestamp()
			}
		}
	}
}
*/
func init() {
	core.ControllersToRegister.PushBack(&MessageController{})
}
