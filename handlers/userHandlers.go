package handlers

import (
	"fmt"
	"log"

	jwt "github.com/dgrijalva/jwt-go"

	"gopkg.in/kataras/iris.v6"

	"github.com/mopAskItAPI/keys"
	"github.com/mopAskItAPI/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func Login(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {

		user := &models.UserCredentials{}

		if err := ctx.ReadJSON(user); err != nil {
			ctx.JSON(iris.StatusBadRequest, err.Error())
			ctx.WriteString(err.Error())
			return
		}

		existingUser := &models.User{}
		err := db.C("users").Find(bson.M{"email": user.Email}).One(&existingUser)

		if (models.User{}) == *existingUser {
			ctx.JSON(iris.StatusBadRequest, "User not exists")
			return
		}
		if user.Password != existingUser.Password {
			ctx.JSON(iris.StatusForbidden, "Combination of email and password is not correct")
			fmt.Println("Error logging in")
			return
		}

		type MyCustomClaims struct {
			jwt.StandardClaims
		}

		claims := MyCustomClaims{
			jwt.StandardClaims{
				ExpiresAt: 15000,
				Issuer:    "user",
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(keys.SignKey)
		fmt.Printf("%v %v", ss, err)

		if err != nil {
			ctx.JSON(iris.StatusInternalServerError, "Error signing token")
			log.Printf("Error signing token: %v\n", err)
		}

		response := models.Token{ss}
		ctx.JSON(iris.StatusOK, response)
	}
}

func GetUsers(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		err, users := models.GetUsers(db)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, "Failed to read request data")
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}
		ctx.JSON(iris.StatusOK, users)
	}
}

func CreateUser(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		user := &models.User{}

		if err := ctx.ReadJSON(user); err != nil {
			ctx.JSON(iris.StatusBadRequest, "Failed to read request data")
			ctx.WriteString(err.Error())
			return
		}

		err := models.CreateUser(db, user)

		if err != nil {
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("User created!")
			ctx.JSON(iris.StatusCreated, "User created")
		}
	}
}

func UpdateUser(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		err, users := models.UpdateUser(db)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, "Failed to read request data")
			ctx.WriteString(err.Error())
		} else {
			fmt.Println("Results All: ", users)
		}
		ctx.JSON(iris.StatusOK, users)
	}
}
