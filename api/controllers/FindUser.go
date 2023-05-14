package controllers

import (
	"net/http"

	"tanzuquiz/initializers"
	"tanzuquiz/models"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get user
func FindUser(c *gin.Context) {
	var user models.Users

	if err := initializers.DB.Where("email = ?", c.Param("email")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
