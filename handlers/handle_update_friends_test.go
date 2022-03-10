package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleUpdateFriendsOK() {
	friends := map[string][]string{
		"friends": {
			"18dd75e9-3d4a-48e2-bafc-3c8f95a8f0d1",
			"f9a9af78-6681-4d7d-8ae7-fc41e7a24d08",
			"2d18862b-b9c3-40f5-803e-5e100a520249",
		},
	}
	body, _ := json.Marshal(friends)
	t.mock.On("UpdateUserFriends", t.userID, friends["friends"]).Return(nil)

	resp := PerformRequest(t.router, "PUT", "/v1/users/"+t.userID+"/friends", string(body))

	assert.Equal(t.T(), http.StatusNoContent, resp.Code)
	assert.Empty(t.T(), resp.Body.String())
}

func (t *HandlersSuiteTest) TestHandlers_HandleUpdateFriendsBadRequest() {
	resp := PerformRequest(t.router, "PUT", "/v1/users/"+t.userID+"/friends", nil)
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "invalid request", respBody["message"])
}

func (t *HandlersSuiteTest) TestHandlers_HandleUpdateFriendsUnprocessableEntity() {
	friends := map[string][]string{
		"friends": {},
	}
	body, _ := json.Marshal(friends)
	t.mock.On("UpdateUserFriends", t.userID, friends["friends"]).Return(errors.New("DB error")).Twice()

	resp := PerformRequest(t.router, "PUT", "/v1/users/"+t.userID+"/friends", string(body))
	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "we couldn't update user friends", respBody["message"])
}
