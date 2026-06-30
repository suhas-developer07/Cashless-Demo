package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(
		"pgx",
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		return nil, fmt.Errorf("Connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Connect: %w", err)
	}
	return db, nil
}
