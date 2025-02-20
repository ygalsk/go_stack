package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("api/health", func (c *gin.Context)  {
		c.JSON(200, gin.H{"status": "API running!"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is runninng on port", port)
	r.Run(":" + port)
}