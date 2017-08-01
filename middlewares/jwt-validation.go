package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/kataras/iris.v6"

	"github.com/dgrijalva/jwt-go"
	"github.com/mopAskItAPI/keys"
)

func ValidateTokenMiddleware(ctx *iris.Context) {
	tokenString := ctx.RequestHeader("Authorization")
	fmt.Printf(tokenString)

	type MyCustomClaims struct {
		Email string `json:"email"`
		jwt.StandardClaims
	}

	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return keys.SignKey, nil
		})

		if err == nil {
			if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
				fmt.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
				ctx.Set("userEmail", claims.Email)
				ctx.Next()
			} else {
				ctx.JSON(http.StatusUnauthorized, "Unauthorised")
				fmt.Printf("Token is not valid")
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, "Unauthorised")
			fmt.Printf("Unauthorised access to this resource")
		}
	})
}

func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
