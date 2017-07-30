package main

import (
	"fmt"

	"github.com/mopAskItAPI/handlers"
	"github.com/mopAskItAPI/models"

	mgo "gopkg.in/mgo.v2"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	database := session.DB("mopAskIt")

	// Regster custom handler for specific http errors.
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx context.Context) {
		// .Values are used to communicate between handlers, middleware.
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal server error: %s", errMessage)
			return
		}

		ctx.Writef("(Unexpected) internal server error")
	})

	app.Use(func(ctx context.Context) {
		ctx.Application().Logger().Infof("Begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	app.Get("/users", userHandlers.GetUsers(database))

	app.Post("/users", func(ctx context.Context) {

		user := &userModel.User{}
		if err := ctx.ReadJSON(user); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		err := database.C("users").Insert(user)
		if err != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("User created!")
			ctx.StatusCode(iris.StatusCreated)
			ctx.WriteString("User created!")
		}
	})
	app.Run(iris.Addr(":8080"))
}
