package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/infrastructure"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

func main() {
	migrator, err := migrate.New(
		"file://db/migrations",
		os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	resources, err := infrastructure.GetResourcesFromPath(os.Getenv("RESOURCES_PATH"))

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := resource.NewPostgresRepository(db)

	for _, resource := range resources {
		err = repository.Save(resource)
		if err != nil {
			log.Fatal(err)
		}
	}
}
