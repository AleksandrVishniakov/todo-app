package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DBConfigs struct {
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *DBConfigs) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
