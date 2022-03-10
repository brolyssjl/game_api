package mocks

import (
	"github.com/brolyssjl/game_api/models"
	"github.com/stretchr/testify/mock"
)

type MockConnection struct {
	mock.Mock
}

func (m *MockConnection) InsertUser(userId, name string) error {
	args := m.Called(mock.Anything, name)

	return args.Error(0)
}

func (m *MockConnection) UpdateUserGameState(gsu models.GameStateUpdate) (int, error) {
	args := m.Called(gsu)

	return args.Int(0), args.Error(1)
}

func (m *MockConnection) InsertUserGameState(gsu models.GameStateUpdate) error {
	args := m.Called(gsu)

	return args.Error(0)
}
