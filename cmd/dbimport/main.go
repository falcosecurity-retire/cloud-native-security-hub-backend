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
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

func main() {
	migrateDatabase()

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	importResources(db)
	importVendors(db)
}

func migrateDatabase() {
	migrator, err := migrate.New("file://db/migrations", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Up()
	if err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}

func importResources(db *sql.DB) {
	resources, err := infrastructure.GetResourcesFromPath(os.Getenv("RESOURCES_PATH"))
	repository := resource.NewPostgresRepository(db)

	for _, resource := range resources {
		err = repository.Save(resource)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func importVendors(db *sql.DB) {
	vendors, err := infrastructure.GetVendorsFromPath(os.Getenv("VENDOR_PATH"))
	repository := vendor.NewPostgresRepository(db)
	for _, vendor := range vendors {
		err = repository.Save(vendor)
		if err != nil {
			log.Fatal(err)
		}
	}
}
