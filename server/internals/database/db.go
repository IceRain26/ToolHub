package database

import (
	"database/sql"
	"fmt"

	"github.com/IceRain26/ToolHub/internals/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(cfg config.DatabaseConfig) error {

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}

	DB = db

	return nil
}
