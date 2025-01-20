package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/TaskFlow/src/presentation/controllers"
)

type Handler struct {
	registrationController *controllers.RegistrationController
 }

func NewHandler(r *controllers.RegistrationController) *Handler {
	return &Handler{
		registrationController: r,
	}
}

func (h *Handler) InitRouter() {
	router := gin.Default()
	
	router.POST("/registration", h.registrationController.Registration)

	router.Run()
 }