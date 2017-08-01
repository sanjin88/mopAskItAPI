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
	User      UserDTO       `json:"user"`
	Responses []Response    `json:"response" bson:",omitempty"`
}

type Response struct {
	Content   string  `json:"content"`
	Votes     []Vote  `json:"votes"`
	CreatedAt string  `json:"created_at"`
	User      UserDTO `json:"user"`
}

type Vote struct {
	Value int     `json:"value"`
	User  UserDTO `json:"user"`
}

func GetQuestions(db *mgo.Database, page int) (error, []Question) {

	questions := []Question{}
	err := db.C("questions").Find(nil).Skip((page - 1) * 20).Limit(20).Sort("-created_at").All(&questions)
	if err != nil {
		panic(err)
	}
	return nil, questions
}

func GetMyQuestions(db *mgo.Database, page int, currentUser User) (error, []Question) {
	fmt.Println(currentUser.Email)

	questions := []Question{}
	err := db.C("questions").Find(bson.M{"user.email": currentUser.Email}).Skip((page - 1) * 20).Limit(20).Sort("-created_at").All(&questions)
	if err != nil {
		panic(err)
	}
	return nil, questions
}

func SaveQuestion(db *mgo.Database, question *Question) error {
	err := db.C("questions").Insert(question)
	return err
}

func VoteQuestion(db *mgo.Database, questionId string, vote *Vote) error {
	question := &Question{}
	err := db.C("questions").FindId(bson.ObjectIdHex(questionId)).One(&question)
	if err != nil {
		panic(err)
	}
	votesInQuestion := question.Votes
	updated := false
	if len(votesInQuestion) < 1 {
		votesInQuestion = append(votesInQuestion, *vote)
		updated = true
	}
	if updated == false {
		for i := 0; i < len(votesInQuestion); i++ {
			if votesInQuestion[i].User.ID == vote.User.ID {
				votesInQuestion[i].Value = vote.Value
				updated = true
			}
		}
	}
	if updated == false {
		votesInQuestion = append(votesInQuestion, *vote)
		updated = true
	}

	err1 := db.C("questions").Update(bson.M{"_id": bson.ObjectIdHex(questionId)}, bson.M{"$set": bson.M{"votes": votesInQuestion}})

	return err1
}

func ResponseOnQuestion(db *mgo.Database, questionId string, response *Response) error {
	question := &Question{}
	err := db.C("questions").FindId(bson.ObjectIdHex(questionId)).One(&question)
	if err != nil {
		panic(err)
	}
	responsesOnQuestion := question.Responses

	responsesOnQuestion = append(responsesOnQuestion, *response)

	err1 := db.C("questions").Update(bson.M{"_id": bson.ObjectIdHex(questionId)}, bson.M{"$set": bson.M{"responses": responsesOnQuestion}})

	return err1
}
