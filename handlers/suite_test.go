package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type header struct {
	Key   string
	Value string
}

type HandlersSuiteTest struct {
	suite.Suite
	router   *gin.Engine
	handlers Handler
	userID   string
	mock     *mocks.MockConnection
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlersSuiteTest))
}

// Setup db value
func (t *HandlersSuiteTest) SetupSuite() {
	log.Println("=== Start Handlers Test Suite Execution ===")

	gin.SetMode(gin.TestMode)
	t.mock = new(mocks.MockConnection)
	t.router = gin.Default()
	engine := engine.NewEngine(t.mock)
	t.handlers = NewHandler(engine)

	// set routes to test
	t.router.POST("/v1/users", t.handlers.HandleCreateUser)
	t.router.PUT("/v1/users/:user_id/states", t.handlers.HandleUpdateGameState)
	t.router.GET("/v1/users/:user_id/states", t.handlers.HandleLoadGameState)
	t.router.PUT("/v1/users/:user_id/friends", t.handlers.HandleUpdateFriends)
	t.router.GET("/v1/users/:user_id/friends", t.handlers.HandleLoadFriends)
}

// Run After All Test Done
func (t *HandlersSuiteTest) TearDownSuite() {
	log.Println("=== End Handlers Test Suite Execution ===")
}

func PerformRequest(r http.Handler, method, path string, payload interface{}, headers ...header) *httptest.ResponseRecorder {
	body := ``
	if payload != nil {
		body = payload.(string)
	}
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
