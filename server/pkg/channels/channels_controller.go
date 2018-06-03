package channels

import (
	"log"

	"github.com/NNeast/talkaneast/server/pkg/messages"

	"github.com/NNeast/talkaneast/server/pkg/core"
	"github.com/fatih/structs"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ChannelsController struct {
	app *core.Application
}

func (cc *ChannelsController) Register(app *core.Application) {
	cc.app = app

	indexErr := app.Db.C("channels").EnsureIndex(mgo.Index{
		Key:    []string{"name"},
		Unique: true,
	})

	if indexErr != nil {
		log.Panicf("Failed to create unique username index in channels controller: %v", indexErr)
	}
}

func (cc *ChannelsController) EventCreateChannel(client *core.Client, session *core.Session, data map[string]interface{}) {
	ch := &Channel{
		Name: data["name"].(string),
	}

	ch.ID = bson.NewObjectId()

	insertErr := cc.app.Db.C("channels").Insert(&ch)

	if insertErr == nil {
		event := core.CreateEvent("ChannelCreated", structs.Map(&ch))
		client.SendEvent(&event)
		log.Printf("event create channel")
	}
}

func (cc *ChannelsController) EventListChannels(client *core.Client, session *core.Session, data map[string]interface{}) {
	var results []Channel
	err := cc.app.Db.C("channels").Find(nil).All(&results)
	if err != nil {
		log.Print(err)
	}
	log.Printf("event list channels")
	event := core.CreateEvent("ListChannels", results)
	client.SendEvent(&event)
}

func (cc *ChannelsController) EventMessage(client *core.Client, session *core.Session, data map[string]interface{}) {
	msg := &messages.Message{
		Content:  data["content"].(string),
		AuthorID: data["author_id"].(string),
	}

	msg.ID = bson.NewObjectId()

	err := cc.app.Db.C("channels").Update(bson.M{
		"_id": bson.ObjectIdHex(data["channelID"].(string)),
	}, bson.M{
		"$push": bson.M{
			"messages": msg,
		},
	})

	if err != nil {
		log.Print(err)
	}
}

func init() {
	core.ControllersToRegister.PushBack(&ChannelsController{})
}
