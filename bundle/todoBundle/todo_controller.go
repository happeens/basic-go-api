package todoBundle

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"

	"github.com/happeens/basic-go-api/app"
	"github.com/happeens/basic-go-api/model"
	"gopkg.in/mgo.v2/bson"
)

type todoController struct{}

func (todoController) Index(c *gin.Context) {
	var result []model.Todo
	err := app.DB().C("todos").Find(nil).All(&result)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (todoController) Show(c *gin.Context) {
	result := model.Todo{}
	id := bson.ObjectIdHex(c.Param("id"))
	app.Log.Debugf("looking for id %v", id)

	err := app.DB().C("todos").FindId(id).One(&result)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

type createRequest struct {
	Description string `json:"description" binding:"required"`
	Done        string `json:"done"`
}

func (todoController) Create(c *gin.Context) {
	var json createRequest
	err := c.BindJSON(&json)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doneBool := false
	if json.Done == "true" {
		doneBool = true
	}

	insert := model.Todo{
		ID:          bson.NewObjectId(),
		Description: json.Description,
		Done:        doneBool,
	}

	err = app.DB().C("todos").Insert(&insert)
	if err != nil {
		app.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": insert.ID})
}

type updateRequest struct {
	Description string `json:"description" binding:"required"`
	Done        string `json:"done"`
}

func (todoController) Update(c *gin.Context) {
	var json updateRequest
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	doneBool := false
	if json.Done == "true" {
		doneBool = true
	}

	id := bson.ObjectIdHex(c.Param("id"))
	update := model.Todo{
		Description: json.Description,
		Done:        doneBool,
	}

	err = app.DB().C("todos").UpdateId(id, update)
	if err != nil {
		app.Log.Errorf("error updating: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": 1})
}

func (todoController) Destroy(c *gin.Context) {
	id := bson.ObjectIdHex(c.Param("id"))
	err := app.DB().C("todos").RemoveId(id)
	if err != nil {
		app.Log.Errorf("error deleting: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"deleted": 1})
}
