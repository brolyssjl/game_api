package handlers

import (
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleCreateUser(c *gin.Context) {
	var payload models.UserCreatePayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	res, err := h.Engine.CreateUser(payload.Name)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "we couldn't create user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   res.UserID,
		"name": res.Name,
	})
}
