package database

import (
	"database/sql"
	"fmt"

	sqlcDb "github.com/gambitier/goblog/database/sqlc"
	"github.com/gambitier/goblog/startup/config"
	_ "github.com/lib/pq"
)

type Database struct {
	*sqlcDb.Queries
	db *sql.DB
}

func NewDatabase(cfg *config.DBConfig) (*Database, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{
		Queries: sqlcDb.New(db),
		db:      db,
	}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
