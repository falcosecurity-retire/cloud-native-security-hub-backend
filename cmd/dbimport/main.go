package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	migrator, err := migrate.New(
		"file://db/migrations",
		os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	if err := migrator.Up(); err != nil {
		log.Fatal(err)
	}
}
