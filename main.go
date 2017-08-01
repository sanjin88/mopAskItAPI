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

	//Init database
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	database := session.DB("mopAskIt")

	//Public API
	app.Post("/users/login", handlers.Login(database))
	app.Post("/users", handlers.CreateUser(database))
	app.Get("/questions/:page", handlers.GetQuestions(database))
	app.Get("/users", handlers.GetUsers(database))

	//Protected API
	app.Put("/users", middlewares.ValidateTokenMiddleware, handlers.UpdateUser(database))
	app.Get("/users/current", middlewares.ValidateTokenMiddleware, handlers.GetCurrentUser(database))
	app.Get("/myquestions/:page", middlewares.ValidateTokenMiddleware, handlers.GetMyQuestions(database))
	app.Post("/questions", middlewares.ValidateTokenMiddleware, handlers.SaveQuestion(database))
	app.Post("/questions/:questionId/vote", middlewares.ValidateTokenMiddleware, handlers.VoteQuestion(database))
	app.Post("/questions/:questionId/response", middlewares.ValidateTokenMiddleware, handlers.ResponseOnQuestion(database))

	app.Listen(":8080")
}
