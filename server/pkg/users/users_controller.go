package users

import (
	"log"

	"github.com/nneast/talkaneast/server/pkg/core"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersController struct {
	app *core.Application
}

func (uc *UsersController) Register(app *core.Application) {
	uc.app = app

	indexErr := app.Db.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	})

	if indexErr != nil {
		log.Panicf("Failed to create unique username index in users controller: %v", indexErr)
	}
}

func (uc *UsersController) EventRegisterUser(client *core.Client, data map[string]interface{}) {
	u := &User{
		Username: data["username"].(string),
		Email:    data["email"].(string),
	}
	u.HashPassword(data["password"].(string))

	// @TODO VALIDATORS
	/*validationError := u.Validate()
	if validationError != nil {
		return validationError
	}*/

	u.ID = bson.NewObjectId()
	uc.app.Db.C("users").Insert(&u)
}

/*
func (uc *UsersController) EventSetNick(user *User, data map[string]interface{}) {
	if containsUsers(uc.um.Users, data["nickname"].(string)) {
		uc.server.sendNotification(user, "this user already exist")
	}
	log.Printf("event set nick")
	user.Username = data["nickname"].(string)
}*/

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}
