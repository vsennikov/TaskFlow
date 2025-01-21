package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vsennikov/TaskFlow/src/services"
)

type LoginController struct {
	userService services.UserServiceInterface
}

func NewLoginController(u services.UserServiceInterface) *LoginController {
	return &LoginController{u}
}

func (l *LoginController) Login(c *gin.Context) {
	var loginModel RegistrationModel
	if err := c.ShouldBindJSON(&loginModel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token, err := l.userService.Login(loginModel.Email, loginModel.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}