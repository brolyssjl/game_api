package main

import (
	"log"
	"time"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/models"
	"github.com/brolyssjl/game_api/server"
	"github.com/joho/godotenv"
)

func main() {
	config := server.Config{
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	loadEnv()
	db := models.NewDatabaseConnection()
	engine := engine.NewEngine(db)
	routes := server.NewRouter(engine)
	srv := server.NewServer(config, routes)

	srv.ListenAndServe()
}

func loadEnv() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}
