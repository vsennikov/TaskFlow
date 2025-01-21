package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vsennikov/TaskFlow/src/services"
)

type TaskController struct {
	userService services.UserServiceInterface
	taskService services.TaskServiceInterface
}

func NewTaskController(u services.UserServiceInterface, 
	t services.TaskServiceInterface) *TaskController {
	return &TaskController{u, t}
}

type TaskModel struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Priority    string    `json:"priority"`
	Status 		string    `json:"status"`
	Due_date    time.Time `json:"due_date"`
}

func (t *TaskController) CreateTask(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty token"})
		return
	}
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}
	userId, err := t.userService.DecodeToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var task TaskModel
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	taskId, err := t.taskService.CreateTask(userId, task.Title, task.Description, 
		task.Category, task.Priority, task.Status, task.Due_date)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"task_id": taskId})
}