package controllers

import (
	"net/http"

	"tanzuquiz/initializers"
	"tanzuquiz/models"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.Users
	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
