package engine_test

import (
	"log"
	"os"
	"testing"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/models"
	"github.com/gin-gonic/gin"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
)

var (
	host     = "0.0.0.0"
	port     = "33060"
	user     = "root"
	pass     = "test123"
	database = "gamedb_test"
)

type EngineSuiteTest struct {
	suite.Suite
	db     models.Spec
	engine *engine.Engine
	userID string
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

	suite.Run(t, new(EngineSuiteTest))
}

// Setup db value
func (t *EngineSuiteTest) SetupSuite() {
	log.Println("=== Start Engine Test Suite Execution ===")
	gin.SetMode(gin.TestMode)

	//t.db = models.NewDatabaseConnection()
	pool, err := dockertest.NewPool("")
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
		t.engine = &engine.Engine{
			DB: t.db,
		}

		return t.db.(*models.Connection).DB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

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
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

// Run After All Test Done
func (t *EngineSuiteTest) TearDownSuite() {
	sqlDB := t.db.(*models.Connection)
	defer sqlDB.DB.Close()

	// Drop Table
	/*for _, val := range getModels() {
		t.db.Migrator().DropTable(val)
	}*/

	log.Println("=== End Engine Test Suite Execution ===")
}
