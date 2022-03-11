package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/brolyssjl/game_api/models"
	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleLoadAllUsersOK() {
	users := &models.UsersDB{
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
	t.mock.On("GetAllUsers").Return(users, nil).Once()

	resp := PerformRequest(t.router, "GET", "/v1/users", nil)

	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusOK, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.NotEmpty(t.T(), respBody["users"])
	assert.Len(t.T(), respBody["users"], 2)
}

func (t *HandlersSuiteTest) TestHandlers_HandleLoadAllUsersUnprocessableEntity() {
	t.mock.On("GetAllUsers").Return(nil, errors.New("DB error")).Once()

	resp := PerformRequest(t.router, "GET", "/v1/users", nil)

	var respBody map[string]string
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "we couldn't retrieve all users", respBody["message"])
}
