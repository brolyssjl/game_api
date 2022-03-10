package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	DB *sql.DB
}

func NewDatabaseConnection() Spec {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		USER, PASS, HOST, PORT, DBNAME,
	))

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &Connection{DB: db}
}
