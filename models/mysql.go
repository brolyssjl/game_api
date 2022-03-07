package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DB *sql.DB
}

func NewConnection() Spec {
	user := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	database := "gamedb"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", user, pass, host, database))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Connection{db: db}
}
