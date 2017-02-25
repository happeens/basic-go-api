package app

import (
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/happeens/basic-go-api/model"
)

type authSettings struct {
	Realm           string
	SigningAlorithm string
	Secret          []byte
	Timeout         time.Duration
	RefreshTimeout  time.Duration
}

var auth authSettings

func initAuth() {
	key := Env("SECRET", "")
	if key == "" {
		Log.Fatal("secret missing from env")
		return
	}

	auth = authSettings{
		Realm:          "test",
		Secret:         []byte(key),
		Timeout:        time.Hour,
		RefreshTimeout: time.Hour,
	}
}

func CreateToken(user model.User) string {
	claims := jwt.MapClaims{
		"user": user.Name,
		"exp":  time.Now().Add(auth.Timeout).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, _ := token.SignedString(auth.Secret)
	return result
}
