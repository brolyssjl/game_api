package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleUpdateGameStateOK() {
	body := `{
		"gamesPlayed": 5,
		"score": 10
	}`
	gameState := models.GameStateUpdate{GamesPlayed: 5, Score: 10, UserId: t.userID}
	t.mock.On("UpdateUserGameState", gameState).Return(1, nil).Once()

	resp := PerformRequest(t.router, "PUT", "/v1/users/"+t.userID+"/states", body)

	assert.Equal(t.T(), http.StatusNoContent, resp.Code)
	assert.Empty(t.T(), resp.Body.String())
}

func (t *HandlersSuiteTest) TestHandlers_HandleUpdateGameStateBadRequest() {
	body := `{
		"gamesPlayed": 5,
		"score": 10
	}`
	gameState := models.GameStateUpdate{GamesPlayed: 5, Score: 10, UserId: t.userID}
	t.mock.On("UpdateUserGameState", gameState).Return(1, nil).Once()

	resp := PerformRequest(t.router, "PUT", "/v1/users/1/states", body)
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "invalid request", respBody["message"])
}
