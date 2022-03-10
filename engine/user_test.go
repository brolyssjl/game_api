package engine_test

import (
	"github.com/brolyssjl/game_api/engine"
	"github.com/stretchr/testify/assert"
)

func (t *EngineSuiteTest) TestEngine_CreateUser() {
	e := &engine.Engine{
		DB: t.db,
	}
	userName := "Jonatan"

	response, err := e.CreateUser(userName)

	assert.NoError(t.T(), err)
	assert.Equal(t.T(), userName, response.Name)
	assert.NotEmpty(t.T(), response.UserID)
}
