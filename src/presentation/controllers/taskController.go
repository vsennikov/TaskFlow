package controllers

import (
	"errors"
	"net/http"

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

func (t *TaskController) CreateTask(c *gin.Context) {
	var task TaskModel
	
	tokenString := c.GetHeader("Authorization")
	userId, err := t.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
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

func (t *TaskController) GetTask(c *gin.Context) {
	var taskId struct {
		TaskId uint `json:"task_id"`
	}
	
	tokenString := c.GetHeader("Authorization")
	userID, err := t.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if err := c.ShouldBindJSON(&taskId); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if taskId.TaskId == 0{
		c.JSON(400, gin.H{"error": "empty task_id"})
		return
	}
	task, err := t.taskService.GetTask(taskId.TaskId, userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, task)
}

func (t *TaskController) GetAllTasks(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	userId, err := t.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	
	tasks, err := t.taskService.GetAllTasks(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}

func (t *TaskController) UpdateTask(c *gin.Context) {
	var updates map[string]interface{}

	tokenString := c.GetHeader("Authorization")
	userID, err := t.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if err = c.ShouldBindJSON(&updates); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	taskID := uint(updates["task_id"].(float64))
	if taskID == 0 {
		c.JSON(400, gin.H{"error": "empty task_id"})
		return
	}
	delete(updates, "task_id")
	err = t.taskService.UpdateTask(taskID, userID, updates)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "task updated"})
}

func (t *TaskController) DeleteTask(c *gin.Context) {
	var taskId struct {
		TaskID uint `json:"task_id"`	
	}

	tokenString := c.GetHeader("Authorization")
	userID, err := t.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if err := c.ShouldBindJSON(&taskId); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if taskId.TaskID == 0 {
		c.JSON(400, gin.H{"error": "empty task_id"})
		return
	}
	err = t.taskService.DeleteTask(taskId.TaskID, userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "task deleted"})
}

func (t *TaskController) GetBySequence(c *gin.Context) {
	var fieldValue map[string]interface{}
	var field string
	var value interface{}

	tokenString := c.GetHeader("Authorization")
	userID, err := t.checkToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if err = c.ShouldBindJSON(&fieldValue); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	if len(fieldValue) != 1 {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}
	for k, v := range fieldValue {
		field = k
		value = v
	}
	tasks, err := t.taskService.GetBySequence(userID, field, value)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}

func (t *TaskController) checkToken(token string) (uint, error) {
	if token == "" {
		return 0, errors.New("empty token")
	}
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}
	userId, err := t.userService.DecodeToken(token)
	if err != nil {
		return 0, errors.New("Unauthorized")
	}
	return userId, nil
}