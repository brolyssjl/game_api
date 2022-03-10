package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockConnection struct {
	mock.Mock
}

func (m *MockConnection) SaveUser(uuid, name string) error {
	args := m.Called(mock.Anything, name)

	return args.Error(0)
}
