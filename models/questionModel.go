package models

import "gopkg.in/mgo.v2/bson"
import "gopkg.in/mgo.v2"

type Question struct {
	ID        bson.ObjectId `json:"id"        bson:"_id,omitempty"`
	Content   string        `json:"content"`
	Votes     []Vote        `json:"votes" bson:",omitempty"`
	CreatedAt string        `json:"created_at"`
	User      User          `json:"user"`
	Responses []Response    `json:"response" bson:",omitempty"`
}

type Response struct {
	Content   string `json:"username"`
	Votes     []Vote `json:"votes"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Vote struct {
	Value int  `json:"value"`
	User  User `json:"user"`
}

func GetQuestions(db *mgo.Database) error {
	questions := []Question{}
	err := db.C("questions").Find(nil).All(&questions)
	return err
}
