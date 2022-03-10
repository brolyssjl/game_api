package engine_test

import (
	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/models"
	"github.com/stretchr/testify/assert"
)

func (t *EngineSuiteTest) TestEngine_CreateUser() {
	e := &engine.Engine{
		DB: t.db,
	}
	userName := "Jonatan"

	response, err := e.CreateUser(userName)
	t.userID = response.UserID

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), userName, response.Name)
	assert.NotEmpty(t.T(), t.userID)
}

func (t *EngineSuiteTest) TestEngine_UpdateUserGameState() {
	e := &engine.Engine{
		DB: t.db,
	}

	gameState := models.GameStatePayload{
		GamesPlayed: 10,
		Score:       100,
	}

	err := e.UpdateUserGameState(t.userID, gameState)

	assert.NoError(t.T(), err)
}
