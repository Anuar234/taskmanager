package handlers

import (
	"net/http"
	"taskmanager/internal/models"
	"taskmanager/internal/mongodb"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect format"})
		return
	}

	if task.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty space"})
		return
	}

	err := mongodb.InsertTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error during the saving of task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}