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
	t.router.PUT("/v1/users/:user_id/state", t.handlers.HandleUpdateGameState)
}

// Run After All Test Done
func (t *HandlersSuiteTest) TearDownSuite() {
	log.Println("=== End Handlers Test Suite Execution ===")
}

func PerformRequest(r http.Handler, method, path, payload string, headers ...header) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(payload))
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
