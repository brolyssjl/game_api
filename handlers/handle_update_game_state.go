package handlers

import (
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleUpdateGameState(c *gin.Context) {
	var request models.GameState
	var user models.UserIDParam
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	err := h.Engine.UpdateUserGameState(user.UserID, models.GameState{
		GamesPlayed: request.GamesPlayed,
		Score:       request.Score,
	})

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "we couldn't update game state",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
