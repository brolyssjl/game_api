package models

type GameStatePayload struct {
	GamesPlayed int `json:"gamesPlayed"`
	Score       int `json:"score"`
}

type GameStateUpdate struct {
	GamesPlayed int
	Score       int
	UserId      string
}

func (db *Connection) UpdateUserGameState(gsu GameStateUpdate) (int, error) {
	res, err := db.DB.Exec("UPDATE game_states SET games_played=?, score=? WHERE user_id=?", gsu.GamesPlayed, gsu.Score, gsu.UserId)
	if err != nil {
		return 0, err
	}

	rows, _ := res.RowsAffected()

	return int(rows), nil
}

func (db *Connection) InsertUserGameState(gsu GameStateUpdate) error {
	_, err := db.DB.Exec("INSERT INTO game_states (games_played, score, user_id) VALUES (?, ?, ?)", gsu.GamesPlayed, gsu.Score, gsu.UserId)
	if err != nil {
		return err
	}

	return nil
}

func (db *Connection) GetUserGameState(userId string) error {
	_, err := db.DB.Exec("SELECT games_played, score FROM game_states WHERE user_id = ?", userId)
	if err != nil {
		return err
	}

	return nil
}
