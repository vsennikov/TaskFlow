package repository

import (
	"time"

	"github.com/vsennikov/TaskFlow/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskDB struct {
}

type TaskDBInterface interface {
	CreateTask(userID uint, title string, description string, category string,
		priority string, status string, dueDate time.Time) (uint, error)
	GetTask(taskID uint, userID uint) (models.TaskJSON, error)
	GetAllTasks(userID uint) ([]models.TaskJSON, error)
	UpdateTask(taskID uint, userID uint, updates map[string]interface{}) error
	DeleteTask(taskID uint, userID uint) error
	GetBySequence(userID uint, field string, value interface{}) ([]models.TaskJSON, error)
}

func NewTaskDB() *TaskDB {
	return &TaskDB{}
}

func (Task) TableName() string {
	return "tasks"
}

func getTaskDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&Task{})
	return db
}

func (*TaskDB) CreateTask(userID uint, title string, description string, category string,
	 priority string, status string, dueDate time.Time) (uint, error) {
	task := Task{UserID: userID, Title: title, Description: description, Category: category,
		 Priority: priority, Status: status, DueDate: dueDate}
	err := getTaskDB().Save(&task).Error
	return task.Model.ID, err
}

func (t *TaskDB) GetTask(taskID uint, userID uint) (models.TaskJSON, error) {
	var task Task
	err := getTaskDB().Where("id = ?", taskID).Where("user_id = ?", userID).First(&task).Error
	return t.transferTask(task), err
}

func (t *TaskDB) GetAllTasks(userID uint) ([]models.TaskJSON, error) {
	var tasks []Task
	var tasksJSON []models.TaskJSON
	err := getTaskDB().Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	for _, task := range tasks {
		tasksJSON = append(tasksJSON, t.transferTask(task))
	}
	return tasksJSON, nil
}

func (t *TaskDB) UpdateTask(taskID uint, userID uint, updates map[string]interface{}) error {
	err := getTaskDB().Model(&Task{}).Where("id = ?", taskID).Where("user_id = ?", userID).Updates(updates).Error
	return err
}

func (*TaskDB) transferTask(task Task) (models.TaskJSON) {
	return models.TaskJSON{TaskID: task.Model.ID, UserID: task.UserID, Title: task.Title,
		Description: task.Description, Category: task.Category, Priority: task.Priority,
		Status: task.Status, DueDate: task.DueDate, CreatedAt: task.CreatedAt, UpdatedAt: task.UpdatedAt}
}

func (*TaskDB) DeleteTask(taskID uint, userID uint) error {
	err := getTaskDB().Where("id = ?", taskID).Where("user_id = ?", userID).Delete(&Task{}).Error
	return err
}

func (t *TaskDB) GetBySequence(userID uint, field string, value interface{}) ([]models.TaskJSON, error) {
		var tasks []Task
		var tasksJSON []models.TaskJSON

		err := getTaskDB().Where("user_id = ?", userID).Where(field + " = ?", value).Find(&tasks).Error
		if err != nil {
			return nil, err
		}
		for _, task := range tasks {
			tasksJSON = append(tasksJSON, t.transferTask(task))
		}
		return tasksJSON, nil
	}
