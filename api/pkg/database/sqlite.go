package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

// Default database path
const DefaultDBPath = "storage/sqlite/playlist.db"

func NewSQLiteDB(dbPath string) (*sql.DB, error) {
	// If no path is provided, use default
	if dbPath == "" {
		dbPath = DefaultDBPath
	}

	// Ensure the directory exists
	err := ensureDir(filepath.Dir(dbPath))
	if err != nil {
		return nil, err
	}

	// Open database connection
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Create songs table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS songs (
            id TEXT PRIMARY KEY,
            title TEXT NOT NULL,
            artist TEXT NOT NULL,
            album TEXT NOT NULL,
            year INTEGER NOT NULL,
            genre TEXT NOT NULL,
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ensureDir creates a directory if it doesn't exist
func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}
