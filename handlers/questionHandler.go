package handlers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/mopAskItAPI/models"
	mgo "gopkg.in/mgo.v2"
)

func GetQuestions(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {

		questions := []models.Question{}

		err := models.GetQuestions(db)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", questions)
		}

		ctx.JSON(questions)
	}
}
func SaveQuestion(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {
		question := &models.Question{}

		if err := ctx.ReadJSON(question); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		err := db.C("questions").Insert(question)
		if err != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Question created!")
			ctx.StatusCode(iris.StatusCreated)
			ctx.WriteString("Question created!")
		}
	}
}
func VoteQuestion(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {

		users := []models.User{}

		err := db.C("users").Find(nil).All(&users)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}

		ctx.JSON(users)
	}
}
func ResponseOnQuestion(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {

		users := []models.User{}

		err := db.C("users").Find(nil).All(&users)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}

		ctx.JSON(users)
	}
}
