package app

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strings"
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
		Log.Critical("secret missing from env")
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

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("WWW-Authenticate", "JWT realm="+auth.Realm)
		if authenticateUser(c) == false {
			c.Abort()
		}
	}
}

func authenticateUser(c *gin.Context) bool {
	token, err := parseHeader(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return false
	}

	claims := token.Claims.(jwt.MapClaims)
	user := claims["user"].(string)

	if user == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user"})
		return false
	}

	c.Set("user", user)
	return true
}

func parseHeader(header string) (*jwt.Token, error) {
	if header == "" {
		return nil, errors.New("empty authorization header")
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid authorization header")
	}

	Log.Debugf("Found token: %v", parts[1])

	return jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("invalid singing algorithm")
		}

		return auth.Secret, nil
	})
}
