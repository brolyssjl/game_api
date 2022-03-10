package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/mocks"
	"github.com/stretchr/testify/assert"
)

func (t *HandlersSuiteTest) TestHandlers_HandleCreateUserOK() {
	body := `{"name": "Jonatan"}`

	resp := PerformRequest(t.router, "POST", "/v1/users", body)

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
	mock := new(mocks.MockConnection)
	mock.On("SaveUser", "uuid", "tester").Return(errors.New("DB error"))
	t.handlers = NewHandler(&engine.Engine{DB: mock})

	resp := PerformRequest(t.router, "POST", "/v1/users", body)

	var respBody map[string]interface{}
	_ = json.Unmarshal(resp.Body.Bytes(), &respBody)

	assert.Equal(t.T(), http.StatusUnprocessableEntity, resp.Code)
	assert.NotEmpty(t.T(), resp.Body.String())
	assert.Equal(t.T(), "we couldn't create user", respBody["message"])
}
