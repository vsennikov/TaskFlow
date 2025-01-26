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

//Should transfer all check from controller to service
//Should also and check for SQLinjections

type TaskServiceInterface interface {
	CreateTask(userID uint, taskName string, taskDescription string, 
		taskCategory string, taskPriority string, taskStatus string, taskDueDate time.Time) (uint, error)
	GetTask (taskId uint, userID uint) (models.TaskJSON, error)
	GetAllTasks(userID uint) ([]models.TaskJSON, error)
	UpdateTask(taskId uint, userID uint, updates map[string]interface{}) error
	DeleteTask(taskID uint, userID uint) error
	GetBySequence(userID uint, field string, value interface{}) ([]models.TaskJSON, error)
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

func (t *TaskService) GetTask(taskId uint, userID uint) (models.TaskJSON, error) {
	return t.repository.GetTask(taskId, userID)
}

func (t *TaskService) GetAllTasks(userID uint) ([]models.TaskJSON, error) {
	return t.repository.GetAllTasks(userID)
}

func (t *TaskService) UpdateTask(taskId uint, userID uint, updates map[string]interface{}) error {
	return t.repository.UpdateTask(taskId, userID, updates)
}

func (t *TaskService) DeleteTask(taskID uint, userID uint) error {
	return t.repository.DeleteTask(taskID, userID)
}

func (t *TaskService) GetBySequence(userID uint, field string, value interface{}) ([]models.TaskJSON, error) {
	return t.repository.GetBySequence(userID, field, value)
}