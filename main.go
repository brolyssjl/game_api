package main

import (
	"log"
	"os"
	"time"

	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/models"
	"github.com/brolyssjl/game_api/server"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	gin.SetMode(os.Getenv("ENVIRONMENT"))

	config := server.Config{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	loadEnv()
	db := models.NewDatabaseConnection()
	engine := engine.NewEngine(db)
	routes := server.NewRouter(engine)
	srv := server.NewServer(config, routes)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func loadEnv() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}
