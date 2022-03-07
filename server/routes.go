package server

import (
	"github.com/brolyssjl/game_api/handlers"
	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health_check", handlers.HandleHealthCheck)

	return router
}
