package main

import (
	"time"

	"github.com/brolyssjl/game_api/models"
	"github.com/brolyssjl/game_api/server"
)

func main() {
	config := models.ServiceConfig{
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	routes := server.NewRouter()
	srv := server.NewServer(config, routes)

	srv.ListenAndServe()
}
