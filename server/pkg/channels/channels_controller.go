package channels

import (
	"log"
	"time"

	"github.com/NNeast/talkaneast/server/pkg/messages"

	"github.com/NNeast/talkaneast/server/pkg/core"
	"github.com/NNeast/talkaneast/server/pkg/users"
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
	err := cc.app.Db.C("channels").Find(nil).Select(bson.M{"messages": 0}).All(&results)
	if err != nil {
		log.Print(err)
	}
	log.Printf("event list channels")
	event := core.CreateEvent("ListChannels", results)
	client.SendEvent(&event)
}

func (cc *ChannelsController) EventMessage(client *core.Client, session *core.Session, data map[string]interface{}) {
	msg := &messages.Message{
		Content:   data["content"].(string),
		AuthorID:  session.UserID,
		CreatedOn: time.Now(),
	}

	msg.ID = bson.NewObjectId()

	err := cc.app.Db.C("channels").Update(bson.M{
		"_id": client.CurrentChannelID,
	}, bson.M{
		"$push": bson.M{
			"messages": msg,
		},
	})

	if err != nil {
		log.Print(err)
	}
	event := core.CreateEvent("Message", msg)

	var clients []*core.Client = cc.app.ClientManager.Clients

	for _, clienti := range clients {
		if &client.CurrentChannelID != nil {
			if client.CurrentChannelID == client.CurrentChannelID {
				clienti.SendEvent(&event)
			}
		}
	}
}

func (cc *ChannelsController) EventQueryMessages(client *core.Client, session *core.Session, data map[string]interface{}) {

	var results Channel

	err := cc.app.Db.C("channels").Find(
		bson.M{"_id": bson.ObjectIdHex(data["channelID"].(string))}).Select(
		bson.M{
			"messages": bson.M{"$slice": -25},
		},
	).One(&results)

	if err != nil {
		log.Print(err)
	}

	event := core.CreateEvent("QueryMessages", results)
	client.SendEvent(&event)
}

func (cc *ChannelsController) EventSetChannel(client *core.Client, session *core.Session, data map[string]interface{}) {
	var results users.User

	err := cc.app.Db.C("users").Find(bson.M{
		"_id": session.UserID,
	}).One(&results)

	if err != nil {
		log.Print(err)
	}

	err2 := cc.app.Db.C("users").Update(bson.M{"_id": results.ID}, bson.M{"$set": bson.M{"currentChannelID": bson.ObjectIdHex(data["channelID"].(string))}})

	if err2 != nil {
		log.Print(err2)
	}

	event := core.CreateEvent("SetChannel", results)
	client.SendEvent(&event)
}

func init() {
	core.ControllersToRegister.PushBack(&ChannelsController{})
}
