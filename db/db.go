package db

import (
	"database/sql"

	_ "modernc.org/sqlite" // importamos assim para dizer que vamos usar implicitamente
)

var DB *sql.DB

func InitDB() {
	newDB, err := sql.Open("sqlite", "api.sql")
	DB = newDB

	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create USERS table: " + err.Error())
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,

		FOREIGN KEY (user_id) REFERENCES Users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create EVENTS table: " + err.Error())
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,

		FOREIGN KEY (event_id) REFERENCES Events(id)
		FOREIGN KEY (user_id) REFERENCES Users(id)
	)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("Could not create REGISTRATIONS table: " + err.Error())
	}
}
