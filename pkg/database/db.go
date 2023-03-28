package database

import (
	"database/sql"
	"fmt"
	"go-boilerplate/pkg/config"
)

func NewDB() (*sql.DB, error) {
	connectionStr := fmt.Sprintf(
		"host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Get().DBHost,
		config.Get().DBPort,
		config.Get().DBUsername,
		config.Get().DBPassword,
		config.Get().DBName,
	)

	return sql.Open(postgresDriver, connectionStr)
}
