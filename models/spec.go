package models

type Spec interface {
	GetAllUsers() error
	CreateUser(name string) error
	GetUserFriends(userId string) error
	GetGameStateByUserId(userId string) error
	SaveUserGameState(state interface{}) error
	UpdateUserFriends(userId string, friends []string) error
}
