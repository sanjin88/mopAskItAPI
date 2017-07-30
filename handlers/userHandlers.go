package handlers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/mopAskItAPI/models"
	mgo "gopkg.in/mgo.v2"
)

func Login(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {
		models.Login(ctx)

				err := 	models.Login(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}
		ctx.JSON(users)
	}
	}
}

func GetUsers(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {
		err := models.GetUsers(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}
		ctx.JSON(users)
	}
}

func CreateUser(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {
		user := &models.User{}

		if err := ctx.ReadJSON(user); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		err := models.CreateUser(db, user)

		if err != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("User created!")
			ctx.StatusCode(iris.StatusCreated)
			ctx.WriteString("User created!")
		}
	}
}

func UpdateUser(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {

		err := models.UpdateUser(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}
		ctx.JSON(users)
	}
}
