package database

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

const (
	postgresDriver = "postgres"
	migrationPath  = "file://./migrations"
)

func initMigration(path string) (*migrate.Migrate, error) {
	db, err := NewDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(path, postgresDriver, driver)
}

func RunLatestMigration() error {
	m, err := initMigration(migrationPath)
	if err != nil {
		return err
	}

	err = m.Up()
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}

func RollbackLastMigration() error {
	m, err := initMigration(migrationPath)
	if err != nil {
		return err
	}

	err = m.Steps(-1)
	if err == migrate.ErrNoChange || err == nil {
		return nil
	}

	return err
}
