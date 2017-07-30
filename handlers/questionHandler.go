package handlers

import (
	"fmt"

	"github.com/mopAskItAPI/models"
	"gopkg.in/kataras/iris.v6"
	mgo "gopkg.in/mgo.v2"
)

func GetQuestions(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		questions := []models.Question{}

		err, questions := models.GetQuestions(db)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
		} else {
			fmt.Println("Results All: ", questions)
		}

		ctx.JSON(iris.StatusOK, questions)
	}
}

func SaveQuestion(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		question := &models.Question{}

		if err := ctx.ReadJSON(question); err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			return
		}

		err := models.SaveQuestion(db, question)
		if err != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Question created!")
			ctx.JSON(iris.StatusCreated, "Question created")
		}
	}
}

func VoteQuestion(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		vote := models.Vote{}

		if err := ctx.ReadJSON(vote); err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			ctx.WriteString(err.Error())
			return
		}
		vote.User.ID = "dadasdasdasdad"
		questionId := ctx.Param("questionId")
		err := models.VoteQuestion(db, questionId, vote)
		if err != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Vote updated!")
			ctx.JSON(iris.StatusCreated, "Vote updated")
			ctx.WriteString("Vote Updated!")
		}
	}
}
func ResponseOnQuestion(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		users := []models.User{}

		err := db.C("users").Find(nil).All(&users)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}

		ctx.JSON(iris.StatusOK, users)
	}
}
