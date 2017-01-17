package controllers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"strconv"

	"api/models"
)

var todoModel models.TodoModel

type TodoController struct{}

func (TodoController) Index(c *gin.Context) {
	var Todos = models.TodoModel{}
	todos := Todos.All()
	c.JSON(http.StatusOK, todos)
}

func (TodoController) Show(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	todo, err := todoModel.Find(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, todo)
}

type createRequest struct {
	Description string `json:"description" binding:"required"`
	Done        string `json:"done"`
}

func (TodoController) Create(c *gin.Context) {
	var json createRequest
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	doneBool := false
	if json.Done == "true" {
		doneBool = true
	}

	id, err := todoModel.New(json.Description, doneBool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

type updateRequest struct {
	Description string `json:"description" binding:"required"`
	Done        string `json:"done"`
}

func (TodoController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var json updateRequest
	err = c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	doneBool := false
	if json.Done == "true" {
		doneBool = true
	}

	var result int64
	result, err = todoModel.Update(uint(id), json.Description, doneBool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated": result})
}

func (TodoController) Destroy(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	result := todoModel.Destroy(uint(id))

	c.JSON(http.StatusOK, gin.H{"deleted": result})
}
