package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "the server is up and running :)",
	})
}
