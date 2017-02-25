package userBundle

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/happeens/basic-go-api/app"
	"github.com/happeens/basic-go-api/model"
	"gopkg.in/mgo.v2/bson"
)

type userController struct{}

type authenticateRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (userController) Authenticate(c *gin.Context) {
	var json authenticateRequest
	err := c.BindJSON(&json)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}

	err = app.DB().C("users").Find(bson.M{"name": json.Name}).One(&user)
	if err != nil {
		app.Log.Errorf("error finding user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json.Password))
	if err != nil {
		app.Log.Errorf("error comparing passwords: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	token := app.CreateToken(user)

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (userController) Index(c *gin.Context) {
	var result []model.User
	err := app.DB().C("users").Find(nil).All(&result)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (userController) Show(c *gin.Context) {
	result := model.User{}
	id := bson.ObjectIdHex(c.Param("id"))

	err := app.DB().C("users").FindId(id).One(&result)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

type createRequest struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (userController) Create(c *gin.Context) {
	var json createRequest
	err := c.BindJSON(&json)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)

	insert := model.User{
		ID:       bson.NewObjectId(),
		Name:     json.Name,
		Password: string(hash[:]),
	}

	err = app.DB().C("users").Insert(&insert)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": insert.ID})
}

func (userController) Destroy(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	err := app.DB().C("users").RemoveId(id)
	if err != nil {
		app.Log.Errorf("error deleting: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"deleted": 1})
}
