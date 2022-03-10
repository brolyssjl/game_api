package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleCreateUserOK() {
	body := `{"name": "Jonatan"}`
	t.mock.On("InsertUser", "uuid", "Jonatan").Return(nil).Once()

	resp := PerformRequest(t.router, "POST", "/v1/users", body)

	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)
	t.userID = respBody["id"].(string)

	assert.Equal(t.T(), http.StatusOK, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
}

func (t *HandlersSuiteTest) TestHandlers_HandleCreateUserBadRequest() {
	body := `{}`

	resp := PerformRequest(t.router, "POST", "/v1/users", body)

	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusBadRequest, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "invalid request", respBody["message"])
}

func (t *HandlersSuiteTest) TestHandlers_HandleCreateUserUnprocessableEntity() {
	body := `{"name": "tester"}`
	t.mock.On("InsertUser", "uuid", "tester").Return(errors.New("DB error")).Once()

	resp := PerformRequest(t.router, "POST", "/v1/users", body)

	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "we couldn't create user", respBody["message"])
}
