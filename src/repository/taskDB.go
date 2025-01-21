package repository

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TaskDB struct {
}

type TaskDBInterface interface {
	CreateTask(userID uint, title string, description string, category string,
		priority string, status string, dueDate time.Time) (uint, error)
}

func NewTaskDB() *TaskDB {
	return &TaskDB{}
}

func (task) TableName() string {
	return "tasks"
}

func getTaskDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(postgresqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&task{})
	return db
}

func (*TaskDB) CreateTask(userID uint, title string, description string, category string,
	 priority string, status string, dueDate time.Time) (uint, error) {
	task := task{UserID: userID, Title: title, Description: description, Category: category,
		 Priority: priority, Status: status, DueDate: dueDate}
	err := getTaskDB().Save(&task).Error
	return task.Model.ID, err
}