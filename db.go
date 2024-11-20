package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import SQLite driver as a side effect
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create tasks table
	createTasksTable(db)
	
	return db
}

func createTasksTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			done BOOLEAN NOT NULL DEFAULT 0	
		);`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create tasks table: %v", err)
	}
}