package models

type Spec interface {
	//GetAllUsers() error
	InsertUser(userId, name string) error
	InsertUserGameState(gsu GameStateUpdate) error
	GetUserGameState(userId string) (*GameStateDB, error)
	UpdateUserGameState(gsu GameStateUpdate) (int, error)
	UpdateUserFriends(userId string, friends []string) error
	/*GetUserFriends(userId string) error*/
}
