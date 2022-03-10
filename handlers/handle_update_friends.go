package handlers

import (
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) HandleUpdateFriends(c *gin.Context) {
	var user models.UserIDParam
	var payload models.UserFriends

	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
			"error":   err.Error(),
		})
		return
	}

	err := h.Engine.UpdateUserFriends(user.UserID, payload.Friends)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "we couldn't update user friends",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
