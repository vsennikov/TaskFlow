package presentation

import "github.com/vsennikov/TaskFlow/src/presentation/controllers"


func InitControllers() {
	registrationController := controllers.RegistrationController{}

	handler := NewHandler(&registrationController)
	handler.InitRouter()
}