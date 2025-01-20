package controllers

import "github.com/gin-gonic/gin"

type RegistrationController struct {
}

func (r *RegistrationController) NewRegistrationController() *RegistrationController {
	return &RegistrationController{}
}

func (r *RegistrationController) Registration(c *gin.Context) {
	// Implement registration logic here
}