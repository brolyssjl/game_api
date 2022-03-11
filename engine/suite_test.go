package engine_test

import (
	"log"
	"testing"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type EngineSuiteTest struct {
	suite.Suite
	mock   *mocks.MockConnection
	engine *engine.Engine
	userID string
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EngineSuiteTest))
}

func (t *EngineSuiteTest) SetupSuite() {
	log.Println("=== Start Engine Test Suite Execution ===")
	gin.SetMode(gin.TestMode)

	t.mock = new(mocks.MockConnection)
	t.engine = &engine.Engine{
		DB: t.mock,
	}
}

func (t *EngineSuiteTest) TearDownSuite() {
	log.Println("=== End Engine Test Suite Execution ===")
}
