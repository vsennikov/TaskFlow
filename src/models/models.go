package models

import "time"

type TaskJSON struct {
	TaskID    uint   `json:"task_id"`
	UserID    uint   `json:"user_id"`
	Title	 string `json:"title"`
	Description string `json:"description"`
	Category string `json:"category"`
	Priority string `json:"priority"`
	Status string `json:"status"`
	DueDate time.Time `json:"due_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}