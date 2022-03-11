package models

type User struct {
	UserID string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
}

type UserCreatePayload struct {
	Name string `json:"name" binding:"required"`
}

type UserIDParam struct {
	UserID string `uri:"user_id" binding:"required,uuid"`
}

type UsersDB struct {
	Users []User
}

type Users struct {
	Users []User `json:"users"`
}

func (db *Connection) InsertUser(userId, name string) error {
	_, err := db.DB.Exec("INSERT INTO users (id, name) VALUES (?, ?)", userId, name)
	if err != nil {
		return err
	}

	return nil
}

func (db *Connection) GetAllUsers() (*UsersDB, error) {
	rows, err := db.DB.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}

	data := &UsersDB{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.UserID,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}

		data.Users = append(data.Users, user)
	}

	return data, nil
}
