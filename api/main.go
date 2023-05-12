package main

import (
	"net/http"
	"tanzuquiz/initializers"

	"tanzuquiz/models"

	"github.com/gin-gonic/gin"
	"github.com/rhjensen79/files2env"
)

func init() {
	files2env.Import("/bindings/tanzuquiz-db")
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	initializers.DB.AutoMigrate(&models.Users{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.Run()
}
