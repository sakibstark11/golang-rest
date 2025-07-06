package database


import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./todos.db")
	if err != nil {
		return nil, err
	}

	schema := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT NOT NULL,
		done BOOLEAN NOT NULL DEFAULT 0
	);`
	_, err = db.Exec(schema)
	return db, err
}
