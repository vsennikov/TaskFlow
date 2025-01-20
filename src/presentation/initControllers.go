package presentation

import (
	"github.com/vsennikov/TaskFlow/src/presentation/controllers"
	"github.com/vsennikov/TaskFlow/src/repository"
	"github.com/vsennikov/TaskFlow/src/services"
)


func InitControllers() {
	db := repository.UserDB{}
	userService := services.NewUserService(&db)
	registrationController := controllers.NewRegistrationController(userService)
	loginController := controllers.NewLoginController(userService)

	handler := NewHandler(registrationController, loginController)
	handler.InitRouter()
}