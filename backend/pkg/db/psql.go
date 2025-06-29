package db

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/sandonemaki/my_read_memo_Go_API/backend/pkg/config"
)

// NewPSQL :
func NewPSQL(c config.Postgres, logger *slog.Logger, applicationName string) *sql.DB {
	// for test and document
	if c.Pseudo {
		return nil
	}
	var err error
	logger.Info("Connecting Postgres...")
	postgresURL := GetPostgresURL(
		c.DBName,
		c.Host,
		c.Port,
		c.User,
		c.Pass,
		c.Sslmode,
		applicationName,
	)
	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		logger.Error("Failed to open postgres", "err", err)
		panic(err)
	}
	if err = db.Ping(); err != nil {
		logger.Error("Failed to ping postgres", "err", err)
		panic(err)
	}

	return db
}

// GetPostgresURL :
func GetPostgresURL(dbName, host, port, user, pass, sslMode, applicationName string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=%s password=%s application_name=%s",
		host, port, user, dbName, sslMode, pass, applicationName,
	)
}
