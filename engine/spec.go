package engine

import "github.com/brolyssjl/game_api/models"

type Spec interface {
	CreateUser(name string) (*models.User, error)
	UpdateUserGameState(userId string, gs models.GameStatePayload) error
}
