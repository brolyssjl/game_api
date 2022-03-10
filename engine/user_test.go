package engine_test

import (
	"github.com/brolyssjl/game_api/models"
	"github.com/stretchr/testify/assert"
)

func (t *EngineSuiteTest) TestEngine_CreateUser() {
	userName := "Jonatan"

	response, err := t.engine.CreateUser(userName)
	t.userID = response.UserID

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), userName, response.Name)
	assert.NotEmpty(t.T(), t.userID)
}

func (t *EngineSuiteTest) TestEngine_UpdateUserGameState() {
	gameState := models.GameState{
		GamesPlayed: 10,
		Score:       100,
	}

	err := t.engine.UpdateUserGameState(t.userID, gameState)

	assert.NoError(t.T(), err)
}

func (t *EngineSuiteTest) TestEngine_LoadUserGameState() {
	gameState, err := t.engine.LoadUserGameState(t.userID)

	assert.NoError(t.T(), err)
	assert.NotEmpty(t.T(), gameState)
}

func (t *EngineSuiteTest) TestEngine_UpdateUserFriends() {
	friends := []string{
		"18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1",
		"f9a9af78-6681-4d7d-8ae7-fc41e7a24d08",
		"2d18862b-b9c3-40f5-803e-5e100a520249",
	}

	err := t.engine.UpdateUserFriends(t.userID, friends)

	assert.NoError(t.T(), err)
}
