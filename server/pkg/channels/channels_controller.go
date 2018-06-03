package channels

import (
	"github.com/NNeast/talkaneast/server/pkg/core"
)

type ChannelsController struct {
	app *core.Application
}

func (cc *ChannelsController) Register(app *core.Application) {
	cc.app = app
}

/*
func (cc *ChannelsController) EventSetChannel(user *users.User, data map[string]interface{}) {
	if containsChannels(cc.app.cm.Channels, data["channel"].(string)) {
		user.CurrentChannel.Name = data["channel"].(string)
	} else {
		cc.app.cm.sendNotification(user, "this channel does not exist")
	}
}

func (cc *ChannelsController) EventCreateChannel(user *users.User, data map[string]interface{}) {
	log.Printf("event create channel")
	if containsChannels(cc.app.cm.Channels, data["name"].(string)) {
		cc.server.sendNotification(user, "this channel already exist")
	}
	cc.server.Channels = append(cc.app.cm.Channels, &Channel{Name: data["name"].(string)})
}

func (cc *ChannelsController) EventListChannels(user *users.User, data map[string]interface{}) {
	cc.app.sendChannelsToUser(user)
}
*/
func init() {
	core.ControllersToRegister.PushBack(&ChannelsController{})
}
