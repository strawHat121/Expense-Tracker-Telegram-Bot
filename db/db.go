package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Fatal("Error in Opening sqlite db", err)
	}

	defer db.Close()

	createTableQuery := `CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		category TEXT NOT NULL,
		amount REAL NOT NULL,
		comment TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`

	_, err = db.Exec(createTableQuery)

	if err != nil {
		log.Fatal("Error in creating the table", err)
	}

	return db

}
