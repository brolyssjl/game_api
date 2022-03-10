package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

var (
	host     = "127.0.0.1"
	port     = "3306"
	user     = "root"
	pass     = "superman123"
	database = "gamedb_test"
)

type header struct {
	Key   string
	Value string
}

type HandlersSuiteTest struct {
	suite.Suite
	db       models.Spec
	router   *gin.Engine
	handlers Handler
}

func TestSuite(t *testing.T) {
	os.Setenv("DB_HOST", host)
	os.Setenv("DB_PORT", port)
	os.Setenv("DB_USER", user)
	os.Setenv("DB_PASS", pass)
	os.Setenv("DB_NAME", database)
	defer os.Unsetenv("DB_HOST")
	defer os.Unsetenv("DB_PORT")
	defer os.Unsetenv("DB_USER")
	defer os.Unsetenv("DB_PASS")
	defer os.Unsetenv("DB_NAME")

	suite.Run(t, new(HandlersSuiteTest))
}

// Setup db value
func (t *HandlersSuiteTest) SetupSuite() {
	log.Println("=== Start Handlers Test Suite Execution ===")

	/*pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=" + pass,
			"MYSQL_DATABASE=" + database,
		},
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306": {
				{HostIP: host, HostPort: port},
			},
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		t.db = models.NewDatabaseConnection()

		return t.db.(*models.Connection).DB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}*/

	gin.SetMode(gin.TestMode)
	t.db = models.NewDatabaseConnection()
	t.router = gin.Default()
	engine := engine.NewEngine(t.db)
	t.handlers = NewHandler(engine)

	// set routes to test
	t.router.POST("/v1/users", t.handlers.HandleCreateUser)

	// run migrations
	/*driver, _ := mysql.WithInstance(t.db.(*models.Connection).DB, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance("file:///models/migrations/create_test_schema.up.sql", "mysql", driver)
	if err != nil {
		log.Fatalf("error running migrations: %s", err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal(err.Error())
	}*/

	// You can't defer this because os.Exit doesn't care for defer
	/*if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}*/
}

// Run After All Test Done
func (t *HandlersSuiteTest) TearDownSuite() {
	sqlDB := t.db.(*models.Connection)
	defer sqlDB.DB.Close()

	// Drop Table
	/*for _, val := range getModels() {
		t.db.Migrator().DropTable(val)
	}*/

	/*if err := t.pool.Purge(t.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}*/

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
