package engine

import "github.com/brolyssjl/game_api/models"

type Spec interface {
	CreateUser(name string) (*models.User, error)
	UpdateUserGameState(userId string, gs models.GameState) error
	LoadUserGameState(userId string) (*models.GameState, error)
	UpdateUserFriends(userId string, friends []string) error
	LoadUserFriends(userId string) (*models.UserFriends, error)
	LoadAllUsers() (*models.Users, error)
}
