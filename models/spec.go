package models

type Spec interface {
	//GetAllUsers() error
	InsertUser(userId, name string) error
	InsertUserGameState(gsu GameStateUpdate) error
	UpdateUserGameState(gsu GameStateUpdate) (int, error)
	GetUserGameState(userId string) (*GameStateDB, error)
	/*GetUserFriends(userId string) error
	UpdateUserFriends(userId string, friends []string) error*/
}
