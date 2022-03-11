package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/brolyssjl/game_api/utils"
	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleLoadFriendsOK() {
	friends := &models.UserFriendsDB{
		Friends: []models.Friend{
			{
				ID:        "18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1",
				Name:      "George",
				Highscore: utils.NewNullInt64(0),
			},
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
	t.mock.On("GetUserFriends", t.userID).Return(friends, nil).Once()

	resp := PerformRequest(t.router, "GET", "/v1/users/"+t.userID+"/friends", nil)
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusOK, resp.Code)
	assert.NotEmpty(t.T(), respBody["friends"])
	assert.Len(t.T(), respBody["friends"], 3)
}

func (t *HandlersSuiteTest) TestHandlers_HandleLoadFriendsBadRequest() {
	resp := PerformRequest(t.router, "GET", "/v1/users/0/friends", nil)
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "invalid request", respBody["message"])
}

func (t *HandlersSuiteTest) TestHandlers_HandleLoadFriendsUnprocessableEntity() {
	t.mock.On("GetUserFriends", t.userID).Return(nil, errors.New("DB error")).Once()

	resp := PerformRequest(t.router, "GET", "/v1/users/"+t.userID+"/friends", nil)
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "we couldn't retrieve user friends data", respBody["message"])
}
