package main

import (
	"log"
	"os"

	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/infrastructure"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

func main() {
	log.Println("Starting database importing job")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	migrateDatabase(db)
	importResources(db)
	importVendors(db)
}

func migrateDatabase(db *sql.DB) {
	log.Println("Applying migrations")

	config := &postgres.Config{}
	driver, err := postgres.WithInstance(db, config)
	if err != nil {
		log.Fatal(err)
	}

	migrator, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Up()
	if err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}

func importResources(db *sql.DB) {
	log.Println("Importing resources")

	resources, err := infrastructure.GetResourcesFromPath(os.Getenv("RESOURCES_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	repository := resource.NewPostgresRepository(db)

	for _, resource := range resources {
		err = repository.Save(resource)
		if err != nil {
			log.Println(err)
		}
	}
}

func importVendors(db *sql.DB) {
	log.Println("Importing vendors")

	vendors, err := infrastructure.GetVendorsFromPath(os.Getenv("VENDOR_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	repository := vendor.NewPostgresRepository(db)
	for _, vendor := range vendors {
		err = repository.Save(vendor)
		if err != nil {
			log.Println(err)
		}
	}
}
