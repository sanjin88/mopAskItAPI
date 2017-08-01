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

type UserDTO struct {
	ID        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
	Email     string        `json:"email"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type Token struct {
	Token string `json:"token"`
}

func GetUsers(db *mgo.Database) (error, []User) {

	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	if err != nil {
		panic(err)
	}
	return nil, users
}

func CreateUser(db *mgo.Database, user *User) error {
	err := db.C("users").Insert(user)
	return err
}

func UpdateUser(db *mgo.Database) (error, []User) {

	users := []User{}
	err := db.C("users").Find(nil).All(&users)
	if err != nil {
		panic(err)
	}
	return nil, users

}
