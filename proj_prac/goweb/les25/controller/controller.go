package controller

import (
	"github.com/gin-gonic/gin"
	"les25/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(c *gin.Context) {

	todoList, err := models.GetAllList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "invalid ID"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
	c.BindJSON(todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, _ := c.Params.Get("id")

	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"id": "delete"})
	}
}
