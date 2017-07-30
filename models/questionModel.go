package models

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

func GetQuestions(db *mgo.Database) (error, []Question) {

	questions := []Question{}
	err := db.C("questions").Find(nil).All(&questions)
	if err != nil {
		panic(err)
	}
	return nil, questions
}

func SaveQuestion(db *mgo.Database, question *Question) error {
	err := db.C("questions").Insert(question)
	return err
}

func VoteQuestion(db *mgo.Database, questionId string, vote Vote) error {
	question := &Question{}
	err := db.C("questions").FindId(bson.ObjectIdHex(questionId)).One(&question)
	if err != nil {
		panic(err)
	}
	votesInQuestion := question.Votes
	updated := false
	if len(votesInQuestion) < 1 {
		votesInQuestion = append(votesInQuestion, vote)
		updated = true
	}
	if updated == false {
		for i := 1; i < len(votesInQuestion); i += 4 {
			if votesInQuestion[i].User.ID == vote.User.ID {
				votesInQuestion[i].Value = vote.Value
				updated = true
			}
		}
	}
	if updated == false {
		votesInQuestion = append(votesInQuestion, vote)
		updated = true
	}

	fmt.Println("Results All: ", question)

	return err
}

//err := c.UpdateId(id, bson.M{"$set": bson.M{"name": "updated name"}})
