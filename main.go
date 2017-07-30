package main

import (
	"github.com/mopAskItAPI/keys"
	"github.com/mopAskItAPI/middlewares"

	"github.com/mopAskItAPI/handlers"

	mgo "gopkg.in/mgo.v2"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())
	keys.InitKeys()

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	database := session.DB("mopAskIt")

	// Regster custom handler for specific http errors.
	/*	app.OnErrorCode(iris.StatusInternalServerError, func(ctx context.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})*/

	/*	app.Use(func(ctx context.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})*/

	app.Post("/users/login", handlers.Login(database))

	app.Post("/users", handlers.CreateUser(database))
	app.Put("/users/{userId: path}", handlers.UpdateUser(database))

	app.Get("/questions", handlers.GetQuestions(database))
	app.Post("/questions", iris.ToHandler(middlewares.ValidateTokenMiddleware), handlers.SaveQuestion(database))

	app.Post("/questions/{questionId: string}/vote", handlers.VoteQuestion(database))
	app.Post("/questions/{questionId: string}/response", handlers.ResponseOnQuestion(database))

	app.Listen(":8080")
}
