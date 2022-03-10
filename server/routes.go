package server

import (
	"github.com/brolyssjl/game_api/engine"
	"github.com/brolyssjl/game_api/handlers"
	"github.com/gin-gonic/gin"
)

func NewRouter(e engine.Spec) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health_check", handlers.HandleHealthCheck)

	handler := handlers.NewHandler(e)
	v1 := router.Group("/v1")
	{
		v1.POST("/users", handler.HandleCreateUser)
		v1.PUT("/users/:user_id/states", handler.HandleUpdateGameState)
		v1.GET("/users/:user_id/states", handler.HandleLoadGameState)
	}

	return router
}
