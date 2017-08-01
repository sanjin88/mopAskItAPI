package handlers

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"

	"github.com/mopAskItAPI/keys"

	jwt "github.com/dgrijalva/jwt-go"

	"gopkg.in/kataras/iris.v6"

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

		err1 := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
		if err1 != nil {
			ctx.JSON(iris.StatusForbidden, "Combination of email and password is not correct")
			fmt.Println("Error logging in", err1)
			return
		}

		type MyCustomClaims struct {
			Email string `json:"email"`
			jwt.StandardClaims
		}

		claims := MyCustomClaims{
			user.Email,
			jwt.StandardClaims{
				ExpiresAt: 6400000,
				Issuer:    "test",
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

func GetCurrentUser(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		userEmail := ctx.Value("userEmail")
		user := &models.User{}
		err := db.C("users").Find(bson.M{"email": userEmail}).One(&user)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, "Failed to read request data")
		} else {
			user.Password = ""
			fmt.Println("Results All: ", user)
		}
		ctx.JSON(iris.StatusOK, user)
	}
}

func CreateUser(db *mgo.Database) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		user := &models.User{}

		if err := ctx.ReadJSON(user); err != nil {
			ctx.JSON(iris.StatusBadRequest, "Failed to read request data")
			return
		}
		password := []byte(user.Password)

		hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		user.Password = string(hashedPassword[:])
		err1 := models.CreateUser(db, user)

		if err1 != nil {
			ctx.JSON(iris.StatusBadRequest, "Failed to read request data")
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
		} else {
			fmt.Println("Results All: ", users)
		}
		ctx.JSON(iris.StatusOK, users)
	}
}
