package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//files2env.Import("/bindings/tanzuquiz-db")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.Run()
}
