package models

import (
	"log"

	"github.com/brolyssjl/game_api/utils"
)

type UserFriendsPayload struct {
	Friends []string `json:"friends" binding:"required"`
}

type Friend struct {
	ID        string          `db:"id" json:"id"`
	Name      string          `db:"name" json:"name"`
	Highscore utils.NullInt64 `db:"score" json:"highscore"`
}

type UserFriendsDB struct {
	Friends []Friend
}

type UserFriends struct {
	Friends []Friend `json:"friends"`
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

func (db *Connection) GetUserFriends(userId string) (*UserFriendsDB, error) {
	query := `
		SELECT uf.friend_id, u.name, gs.score
		FROM user_has_friends uf
		LEFT JOIN users u ON u.id = uf.friend_id
		LEFT JOIN game_states gs ON gs.user_id = uf.friend_id
		WHERE uf.user_id = ?`

	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return nil, err
	}

	data := &UserFriendsDB{}
	for rows.Next() {
		friend := Friend{}
		err = rows.Scan(
			&friend.ID,
			&friend.Name,
			&friend.Highscore,
		)
		if err != nil {
			return nil, err
		}

		data.Friends = append(data.Friends, friend)
	}

	return data, nil
}
