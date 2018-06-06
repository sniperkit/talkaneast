package core

import (
	"flag"
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

type Application struct {
	ClientManager *ClientManager
	EventHandler  *EventHandler

	Db         *mgo.Database
	DbOverride string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *Application) handleWS(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	app.ClientManager.registerConnection(c)
	defer c.Close()

	for {
		var response map[string]interface{}
		err := c.ReadJSON(&response)
		if err != nil {
			app.ClientManager.removeConnection(c)
			break
		}

		app.EventHandler.handleEvent(app.ClientManager.getUserByConn(c), response)
	}
}

func (app *Application) Run() {
	app.EventHandler = &EventHandler{
		App: app,
	}

	app.ClientManager = &ClientManager{
		Clients: make([]*Client, 0),
	}

	flag.Parse()
	http.HandleFunc("/ws", app.handleWS)

	log.Info("Connecting to database")
	var dbError error
	sess, dbError := mgo.Dial("mongodb://127.0.0.1:27017")
	if dbError != nil {
		log.Fatal("Failed to connect to database: ", dbError)
		return
	}

	app.Db = sess.DB(app.DbOverride) // get DB specified in DbOverride or the mongo url (https://godoc.org/gopkg.in/mgo.v2#Session.DB)

	app.EventHandler.registerHandlers()
	log.Info("Started listening")
	log.Fatal(http.ListenAndServe(":2148", nil))
}
