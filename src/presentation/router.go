package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/TaskFlow/src/presentation/controllers"
)

type Handler struct {
	registrationController *controllers.RegistrationController
	loginController *controllers.LoginController
}

func NewHandler(r *controllers.RegistrationController, l *controllers.LoginController) *Handler {
	return &Handler{
		registrationController: r,
		loginController: l,
	}
}

func (h *Handler) InitRouter() {
	router := gin.Default()
	
	router.POST("/registration", h.registrationController.Registration)
	router.POST("/login", h.loginController.Login)

	router.Run()
}