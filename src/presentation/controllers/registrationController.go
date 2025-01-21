package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/TaskFlow/src/services"
)

type RegistrationModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RegistrationController struct {
	userService services.UserServiceInterface
}

func NewRegistrationController(u services.UserServiceInterface) *RegistrationController {
	return &RegistrationController{u}
}

func (r *RegistrationController) Registration(c *gin.Context) {
	var userData RegistrationModel
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID, err := r.userService.UserRegistration(userData.Username, userData.Email, userData.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"user_id": userID})
}

