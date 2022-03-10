package handlers

import (
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleLoadGameState(c *gin.Context) {
	var user models.UserIDParam
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	response, err := h.Engine.LoadUserGameState(user.UserID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "we couldn't load game state :(",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
