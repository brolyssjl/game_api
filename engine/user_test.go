package engine_test

import (
	"github.com/brolyssjl/game_api/models"
	"github.com/brolyssjl/game_api/utils"
	"github.com/stretchr/testify/assert"
)

func (t *EngineSuiteTest) TestEngine_CreateUser() {
	userName := "Jonatan"
	t.mock.On("InsertUser", "uuid", userName).Return(nil).Once()

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
	gameStateUpdate := models.GameStateUpdate{
		GamesPlayed: 10,
		Score:       100,
		UserId:      t.userID,
	}
	t.mock.On("UpdateUserGameState", gameStateUpdate).Return(1, nil).Once()

	err := t.engine.UpdateUserGameState(t.userID, gameState)

	assert.NoError(t.T(), err)
}

func (t *EngineSuiteTest) TestEngine_LoadUserGameState() {
	gameStateDB := &models.GameStateDB{
		GamesPlayed: 10,
		Score:       100,
	}
	t.mock.On("GetUserGameState", t.userID).Return(gameStateDB, nil).Once()

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
	t.mock.On("UpdateUserFriends", t.userID, friends).Return(nil)

	err := t.engine.UpdateUserFriends(t.userID, friends)

	assert.NoError(t.T(), err)
}

func (t *EngineSuiteTest) TestEngine_LoadUserFriends() {
	friendsDB := &models.UserFriendsDB{
		Friends: []models.Friend{
			{
				ID:        "2d18862b-b9c3-40f5-803e-5e100a520249",
				Name:      "Doom Slayer",
				Highscore: utils.NewNullInt64(1000),
			},
			{
				ID:        "f9a9af78-6681-4d7d-8ae7-fc41e7a24d08",
				Name:      "Kirito",
				Highscore: utils.NewNullInt64(500),
			},
		},
	}
	t.mock.On("GetUserFriends", t.userID).Return(friendsDB, nil).Once()

	friends, err := t.engine.LoadUserFriends(t.userID)

	assert.NoError(t.T(), err)
	assert.NotEmpty(t.T(), friends)
}

func (t *EngineSuiteTest) TestEngine_LoadAllUsers() {
	usersDB := &models.UsersDB{
		Users: []models.User{
			{
				UserID: "18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1",
				Name:   "George",
			},
			{
				UserID: "2d18862b-b9c3-40f5-803e-5e100a520249",
				Name:   "Doom Slayer",
			},
		},
	}
	t.mock.On("GetAllUsers").Return(usersDB, nil).Once()

	users, err := t.engine.LoadAllUsers()

	assert.NoError(t.T(), err)
	assert.NotEmpty(t.T(), users)
}
