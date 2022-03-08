package main

import (
	"time"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/models"
	"github.com/brolyssjl/game_api/server"
)

func main() {
	config := server.Config{
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	db := models.NewDatabaseConnection()
	engine := engine.NewEngine(db)
	routes := server.NewRouter(engine)
	srv := server.NewServer(config, routes)

	srv.ListenAndServe()
}
