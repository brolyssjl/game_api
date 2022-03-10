package models

type User struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

type UserCreatePayload struct {
	Name string `json:"name" binding:"required"`
}

type UserIDParam struct {
	UserID string `uri:"user_id" binding:"required,uuid"`
}

func (db *Connection) InsertUser(userId, name string) error {
	_, err := db.DB.Exec("INSERT INTO users (id, name) VALUES (?, ?)", userId, name)
	if err != nil {
		return err
	}

	return nil
}
