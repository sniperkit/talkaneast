package users

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username         string        `json:"username" structs:"username" bson:"username"`
	CurrentChannelID bson.ObjectId `json:"currentChannelID" structs:"currentChannelID" bson:"currentChannelID,omitempty"`
	Password         string        `json:"-" bson:"password"`
	Email            string        `json:"email" bson:"email"`
	Avatar           string        `json:"avatar" bson:"avatar"`
}

/*
HashPassword takes a cleartext password and sets the Password field with a hashed version.
*/
func (u *User) HashPassword(password string) error {

	rawHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return errors.Wrap(err, "failed to hash user password:")
	}
	u.Password = string(rawHash)
	return nil
}
