package presentation

import (
	"github.com/vsennikov/TaskFlow/src/presentation/controllers"
	"github.com/vsennikov/TaskFlow/src/repository"
	"github.com/vsennikov/TaskFlow/src/services"
)


func InitControllers() {
	userDb := repository.UserDB{}
	taskDb := repository.TaskDB{}
	userService := services.NewUserService(&userDb)
	taskService := services.NewTaskService(&taskDb)
	registrationController := controllers.NewRegistrationController(userService)
	loginController := controllers.NewLoginController(userService)
	taskController := controllers.NewTaskController(userService, taskService)
	handler := NewHandler(registrationController, loginController, taskController)
	handler.InitRouter()
}