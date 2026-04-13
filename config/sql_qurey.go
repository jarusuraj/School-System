package config

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func AllUPQueryMigration() {
	m, err := migrate.New(
		"file://database/migrations",
		DBUrl(),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
func OneUPQueryMigration() {
	m, err := migrate.New(
		"file://database/migrations",
		DBUrl(),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
func OneDownQueryMigration() {
	m, err := migrate.New(
		"file://database/migrations",
		DBUrl(),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Steps(-1)
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
