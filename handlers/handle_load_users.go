package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleLoadUsers(c *gin.Context) {
	response, err := h.Engine.LoadAllUsers()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "we couldn't retrieve all users",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
