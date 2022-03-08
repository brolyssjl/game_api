package engine

import (
	"github.com/brolyssjl/game_api/models"
	"github.com/google/uuid"
)

func (e *Engine) CreateUser(name string) (*models.User, error) {
	userId := uuid.NewString()

	err := e.DB.SaveUser(userId, name)
	if err != nil {
		return nil, err
	}

	return &models.User{
		UserID: userId,
		Name:   name,
	}, nil
}
