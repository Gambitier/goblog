package database

import (
	"database/sql"
	"fmt"

	sqlcDb "github.com/gambitier/goblog/database/sqlc"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlcDb.Queries
	db *sql.DB
}

func NewDatabase(databaseURL string) (*Database, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &Database{
		Queries: sqlcDb.New(db),
		db:      db,
	}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
