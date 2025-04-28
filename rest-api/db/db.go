package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	//_ "modernc.org/sqlite"
)

// Globally, we will have a pointer to the database handler
var DB *sql.DB

// InitDB is used to make a connection to database
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database")
	}

	// Test connection
	if err = DB.Ping(); err != nil {
		panic(err)
	}

	//Configure connection pools to set up how many simultaneous open connections
	//we want to the database
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	//Create the Tables
	createTables()
}

func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
)`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
CREATE TABLE IF NOT EXISTS events (
 id INTEGER PRIMARY KEY AUTOINCREMENT,
 name TEXT NOT NULL,
 description TEXT NOT NULL,
 location TEXT NOT NULL,
 dateTime DATETIME NOT NULL,
 user_id INTEGER,
FOREIGN KEY(user_id) REFERENCES users(id)
)
`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationsTable := `
CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(event_id) REFERENCES events(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(err)
	}
}
