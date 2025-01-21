package services

import (
	"errors"
	"time"

	"github.com/vsennikov/TaskFlow/src/models"
	"github.com/vsennikov/TaskFlow/src/repository"
)

type TaskService struct {
	repository repository.TaskDBInterface
}

type TaskServiceInterface interface {
	CreateTask(userID uint, taskName string, taskDescription string, 
		taskCategory string, taskPriority string, taskStatus string, taskDueDate time.Time) (uint, error)
	GetTask (taskId uint) (models.TaskJSON, error)
	GetAllTasks(userID uint) ([]models.TaskJSON, error)
}

func NewTaskService(r repository.TaskDBInterface) *TaskService {
	return &TaskService{r}
}

func (t *TaskService) CreateTask(userID uint, taskName string, taskDescription string, 
	taskCategory string, taskPriority string, taskStatus string, taskDueDate time.Time) (uint, error) {
	if taskName == "" {
		return 0, errors.New("task name cannot be empty")
	}
	if containsSQLInjection(taskName) || containsSQLInjection(taskDescription) || 
		containsSQLInjection(taskCategory) || containsSQLInjection(taskPriority) {
		return 0, errors.New("input contains invalid characters")
	}
	if taskStatus == "" {
		taskStatus = "New"
	}
	return t.repository.CreateTask(userID, taskName, taskDescription, taskCategory, taskPriority, taskStatus, taskDueDate)
}

func (t *TaskService) GetTask(taskId uint) (models.TaskJSON, error) {
	return t.repository.GetTask(taskId)
}

func (t *TaskService) GetAllTasks(userID uint) ([]models.TaskJSON, error) {
	return t.repository.GetAllTasks(userID)
}