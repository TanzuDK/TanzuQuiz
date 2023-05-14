package controllers

import (
	"net/http"
	"tanzuquiz/initializers"
	"tanzuquiz/models"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Company string `json:"company" binding:"required"`
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.Users{Name: input.Name, Email: input.Email, Company: input.Company}
	initializers.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
