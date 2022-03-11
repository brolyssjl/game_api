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

func (m *MockConnection) GetUserGameState(userId string) (*models.GameStateDB, error) {
	args := m.Called(userId)

	if args.Get(0) != nil {
		return args.Get(0).(*models.GameStateDB), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockConnection) UpdateUserFriends(userId string, friends []string) error {
	args := m.Called(userId, friends)

	return args.Error(0)
}

func (m *MockConnection) GetUserFriends(userId string) (*models.UserFriendsDB, error) {
	args := m.Called(userId)

	if args.Get(0) != nil {
		return args.Get(0).(*models.UserFriendsDB), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockConnection) GetAllUsers() (*models.UsersDB, error) {
	args := m.Called()

	if args.Get(0) != nil {
		return args.Get(0).(*models.UsersDB), args.Error(1)
	}

	return nil, args.Error(1)
}
