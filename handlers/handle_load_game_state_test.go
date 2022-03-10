package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleLoadUserGameStateOK() {
	gameState := &models.GameStateDB{
		GamesPlayed: 10,
		Score:       100,
	}
	t.mock.On("GetUserGameState", t.userID).Return(gameState, nil).Once()

	resp := PerformRequest(t.router, "GET", "/v1/users/"+t.userID+"/states", nil)

	var respBody map[string]int
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusOK, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), gameState.GamesPlayed, respBody["gamesPlayed"])
	assert.Equal(t.T(), gameState.Score, respBody["score"])
}

func (t *HandlersSuiteTest) TestHandlers_HandleLoadUserGameStateBadRequest() {
	resp := PerformRequest(t.router, "GET", "/v1/users/0/states", nil)

	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "invalid request", respBody["message"])
}

func (t *HandlersSuiteTest) TestHandlers_HandleLoadUserGameStateUnprocessableEntity() {
	t.mock.On("GetUserGameState", t.userID).Return(nil, errors.New("DB error")).Once()

	resp := PerformRequest(t.router, "GET", "/v1/users/"+t.userID+"/states", nil)

	var respBody map[string]string
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "we couldn't load game state :(", respBody["message"])
}
