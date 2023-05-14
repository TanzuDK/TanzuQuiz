package main

import (
	"log"
	"net/http"
	"tanzuquiz/controllers"
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
	server := gin.Default()
	initializers.DB.AutoMigrate(&models.Users{})

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/users", controllers.FindUsers)
	router.GET("/users/:email", controllers.FindUser)
	router.POST("/users", controllers.CreateUser)

	log.Fatal(server.Run())
}
