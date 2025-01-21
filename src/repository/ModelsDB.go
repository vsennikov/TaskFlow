package repository

import (
	"time"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Username string `gorm:"column:username"`
	Email    string `gorm:"column:email;unique"`
	Password string `gorm:"column:password"`
}

type Task struct {
	gorm.Model
	UserID    uint   `gorm:"column:user_id"`
	Title	 string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Category string `gorm:"column:category"`
	Priority string `gorm:"column:priority"`
	Status string `gorm:"column:status"`
	DueDate time.Time `gorm:"column:due_date"`
}