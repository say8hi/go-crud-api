package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"crud_api/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUsersTable() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS Users (
            id SERIAL PRIMARY KEY,
            username TEXT NOT NULL,
            full_name TEXT
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertUser(user models.User) error {
	_, err := db.Exec("INSERT INTO users (username, full_name) VALUES ($1, $2)", user.Username, user.FullName)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id int) (models.User, error) {
	var user models.User
	row := db.QueryRow("SELECT id, username, full_name FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Username, &user.FullName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(id int, user models.User) error {
	_, err := db.Exec("UPDATE users SET username = $1, full_name = $2 WHERE id = $3", user.Username, user.FullName, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
