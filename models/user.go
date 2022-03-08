package models

type User struct {
	Name   string `json:"name" binding:"required"`
	UserID string `json:"user_id" uri:"user_id" binding:"required,uuid"`
}

type UserCreatePayload struct {
	Name string `json:"name" binding:"required"`
}

func (db *Connection) SaveUser(uuid, name string) error {
	_, err := db.DB.Exec("INSERT INTO users (id, name) VALUES (?, ?)", uuid, name)
	if err != nil {
		return err
	}

	return nil
}
