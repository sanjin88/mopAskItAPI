package userHandlers

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/mopAskItAPI/models"
	mgo "gopkg.in/mgo.v2"
)

func GetUsers(db *mgo.Database) context.Handler {
	return func(ctx context.Context) {

		users := []userModel.User{}

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
