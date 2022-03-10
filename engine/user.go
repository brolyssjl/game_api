package engine

import (
	"github.com/brolyssjl/game_api/models"
	"github.com/google/uuid"
)

func (e *Engine) CreateUser(name string) (*models.User, error) {
	userId := uuid.NewString()

	err := e.DB.InsertUser(userId, name)
	if err != nil {
		return nil, err
	}

	return &models.User{
		UserID: userId,
		Name:   name,
	}, nil
}

func (e *Engine) UpdateUserGameState(userId string, gs models.GameState) error {
	gameData := models.GameStateUpdate{
		GamesPlayed: gs.GamesPlayed,
		Score:       gs.Score,
		UserId:      userId,
	}

	rows, err := e.DB.UpdateUserGameState(gameData)
	if err != nil {
		return err
	}

	// user game state not exist
	if rows == 0 {
		err := e.DB.InsertUserGameState(gameData)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Engine) LoadUserGameState(userId string) (*models.GameState, error) {
	data, err := e.DB.GetUserGameState(userId)
	if err != nil {
		return nil, err
	}

	return &models.GameState{
		GamesPlayed: data.GamesPlayed,
		Score:       data.Score,
	}, nil
}
