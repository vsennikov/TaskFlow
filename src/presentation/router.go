package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/TaskFlow/src/presentation/controllers"
)

type Handler struct {
	registrationController *controllers.RegistrationController
	loginController *controllers.LoginController
	taskController *controllers.TaskController
}

func NewHandler(r *controllers.RegistrationController, l *controllers.LoginController,
	t *controllers.TaskController) *Handler {
	return &Handler{
		registrationController: r,
		loginController: l,
		taskController: t,
	}
}

func (h *Handler) InitRouter() {
	router := gin.Default()	
	router.POST("/registration", h.registrationController.Registration)
	router.POST("/login", h.loginController.Login)
	router.POST("/task", h.taskController.CreateTask)
	router.GET("/task", h.taskController.GetTask)
	router.GET("/tasks", h.taskController.GetAllTasks)
	router.PUT("/task", h.taskController.UpdateTask)
	router.DELETE("/task", h.taskController.DeleteTask)
	router.GET("/tasks/sequence", h.taskController.GetBySequence)
	router.Run()
}