package models

import "log"

type UserFriends struct {
	Friends []string `json:"friends" binding:"required"`
}

type Friend struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Highscore int    `json:"highscore"`
}

func (db *Connection) UpdateUserFriends(userId string, friends []string) error {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM user_has_friends WHERE user_id = ?", userId)
	if err != nil {
		return err
	}

	for _, friend := range friends {
		log.Printf("user: %s - friend: %s", userId, friend)
		_, err = tx.Exec("INSERT INTO user_has_friends (user_id, friend_id) VALUES (?, ?)", userId, friend)
		if err != nil {
			log.Printf("%+v", err)
			return err
		}
	}

	return tx.Commit()
}
