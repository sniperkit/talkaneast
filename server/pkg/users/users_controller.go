package users

import (
	"log"
	"time"

	"github.com/NNeast/talkaneast/server/pkg/core"
	"github.com/fatih/structs"
	"golang.org/x/crypto/bcrypt"
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
	insertErr := uc.app.Db.C("users").Insert(&u)

	if insertErr == nil {
		event := core.CreateEvent("UserRegistered", structs.Map(&u))
		client.SendEvent(&event)
	}

}

func (uc *UsersController) EventLoginUser(client *core.Client, data map[string]interface{}) {
	user := &User{}
	findErr := uc.app.Db.C("users").Find(bson.M{
		"username": data["username"].(string),
	}).One(user)
	if findErr != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"].(string))) != nil {
		//core.NewErrorResponse("Invalid username or password.", 401)
	}
	sess := &core.Session{
		ID:        bson.NewObjectId(),
		UserID:    user.ID,
		IPAddress: client.Conn.RemoteAddr().String(),
		Active:    true,
		CreatedOn: time.Now(),
		ExpiresOn: time.Now().Add(time.Hour * 24 * 7),
	}
	sess.AssignToken()
	insertErr := uc.app.Db.C("sessions").Insert(sess)
	if insertErr == nil {
		sessMap := structs.Map(&sess)
		sessMap["user"] = user
		event := core.CreateEvent("SessionData", sessMap)
		client.SendEvent(&event)
	}

	client.UserID = user.ID
}

func (uc *UsersController) EventLogoutUser(client *core.Client, sess *core.Session, data map[string]interface{}) {
	sess.Active = false
	uc.app.Db.C("sessions").Update(bson.M{
		"_id": sess.ID,
	}, bson.M{
		"$set": bson.M{
			"active": false,
		},
	})
}

func (uc *UsersController) EventSetAvatar(client *core.Client, sess *core.Session, data map[string]interface{}) {

	err2 := uc.app.Db.C("users").Update(bson.M{
		"_id": sess.UserID}, bson.M{
		"$set": bson.M{
			"avatar_url": data["avatar_url"].(string)}})

	if err2 != nil {
		log.Print(err2)
	}

}

func init() {
	core.ControllersToRegister.PushBack(&UsersController{})
}
