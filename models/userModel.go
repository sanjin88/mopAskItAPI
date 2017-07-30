package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
}

func GetUsers(db *mgo.Database) error {

	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	return err
}

func Login(db *mgo.Database) error {

	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	return err
}

func CreateUser(db *mgo.Database, user User) error {
	err := db.C("users").Insert(user)
	return err
}

func UpdateUser(db *mgo.Database) error {

	users := []User{}

	err := db.C("users").Find(nil).All(&users)
	return err

}
