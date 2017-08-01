package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/mopAskItAPI/models"
	"gopkg.in/kataras/iris.v6"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func GetQuestions(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		page, err := strconv.ParseInt(ctx.Param("page"), 10, 64)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
		}
		questions := []models.Question{}

		err1, questions := models.GetQuestions(db, int(page))
		if err1 != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
		} else {
			fmt.Println("Results All: ", questions)
		}

		ctx.JSON(iris.StatusOK, questions)
	}
}

func GetMyQuestions(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		userEmail := ctx.Value("userEmail")
		user := &models.User{}
		err := db.C("users").Find(bson.M{"email": userEmail}).One(&user)

		page, err := strconv.ParseInt(ctx.Param("page"), 10, 64)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
		}
		questions := []models.Question{}

		err1, questions := models.GetMyQuestions(db, int(page), *user)
		if err1 != nil {
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
		userEmail := ctx.Value("userEmail")
		user := &models.UserDTO{}
		err := db.C("users").Find(bson.M{"email": userEmail}).One(&user)
		t := time.Now()
		question.CreatedAt = t.Format(time.RFC822)
		question.User = *user
		if err := ctx.ReadJSON(question); err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			return
		}

		err1 := models.SaveQuestion(db, question)
		if err1 != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Question created!")
			ctx.JSON(iris.StatusCreated, "Question created")
		}
	}
}

func VoteQuestion(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		userEmail := ctx.Value("userEmail")
		vote := &models.Vote{}
		user := &models.UserDTO{}
		err := db.C("users").Find(bson.M{"email": userEmail}).One(&user)
		if err != nil {
			panic(err)
		}

		if err := ctx.ReadJSON(&vote); err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			return
		}
		vote.User.ID = user.ID
		vote.User.Email = user.Email
		vote.User.Firstname = user.Firstname
		questionId := ctx.Param("questionId")
		err1 := models.VoteQuestion(db, questionId, vote)
		if err1 != nil {
			ctx.JSON(iris.StatusBadRequest, "Error")
		} else {
			fmt.Println("Vote updated!")
			ctx.JSON(iris.StatusCreated, "Vote updated")
		}
	}
}

func ResponseOnQuestion(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		userEmail := ctx.Value("userEmail")
		response := &models.Response{}
		user := &models.UserDTO{}
		err := db.C("users").Find(bson.M{"email": userEmail}).One(&user)
		if err != nil {
			panic(err)
		}

		if err := ctx.ReadJSON(&response); err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			return
		}
		response.User.ID = user.ID
		response.User.Email = user.Email
		response.User.Firstname = user.Firstname
		t := time.Now()
		response.CreatedAt = t.Format(time.RFC822)
		questionId := ctx.Param("questionId")
		err1 := models.ResponseOnQuestion(db, questionId, response)
		if err1 != nil {
			ctx.JSON(iris.StatusBadRequest, "Error")
		} else {
			fmt.Println("Response Created!")
			ctx.JSON(iris.StatusCreated, "Response Created")
		}
	}
}
