package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaodddev/apex-f1-api/internal/clients"
)

func GetDrivers(c *gin.Context) {
	data, err := clients.GetDrivers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "erro ao buscar pilotos",
		})
		return
	}

	c.Data(200, "application/json", []byte(data))
}
