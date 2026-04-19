package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaodddev/apex-f1-api/internal/handlers"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "online"})
	})

	r.GET("/drivers", handlers.GetDrivers)

	r.Run(":8080")
}
