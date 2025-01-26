package controllers

import "time"

type RegistrationModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type TaskModel struct {
	TaskId	  	uint      `json:"task_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Priority    string    `json:"priority"`
	Status 		string    `json:"status"`
	Due_date    time.Time `json:"due_date"`
}
