package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dataSource string) {
	migration, err := migrate.New(
		"file://migrations",
		dataSource,
	)
	if err != nil {
		log.Fatalf("cannot create migrate instance: %v", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migration completed")
}
