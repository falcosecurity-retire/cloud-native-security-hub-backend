package resource_test

import (
	"database/sql"
	"os"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAResource(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := resource.NewPostgresRepository(db)

	repository.Save(resources.Apache())

	retrieved, _ := repository.FindById("apache")
	assert.Equal(t, resources.Apache(), retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
	db.Exec("TRUNCATE TABLE latest_security_resources")
}

func TestFindAllResources(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := resource.NewPostgresRepository(db)

	repository.Save(resources.Apache())
	repository.Save(resources.MongoDB())

	retrieved, _ := repository.FindAll()
	assert.Equal(t, []*resource.Resource{resources.Apache(), resources.MongoDB()}, retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
	db.Exec("TRUNCATE TABLE latest_security_resources")
}

func TestFindResourceByIdDoesntFindTheResource(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := resource.NewPostgresRepository(db)

	retrieved, err := repository.FindById("non existent id")

	assert.Nil(t, retrieved)
	assert.Equal(t, resource.ErrResourceNotFound, err)
}

func TestFindResourceByIdReturnsLatestVersion(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := resource.NewPostgresRepository(db)

	apache := resources.Apache()
	repository.Save(apache)

	apache.Version = "2.0.0"
	repository.Save(apache)

	retrieved, _ := repository.FindById("apache")

	assert.Equal(t, apache, retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
	db.Exec("TRUNCATE TABLE latest_security_resources")
}

func TestFindResourceByVersionReturnsItsVersion(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := resource.NewPostgresRepository(db)

	apache := resources.Apache()
	repository.Save(apache)

	apache.Version = "2.0.0"
	repository.Save(apache)

	retrieved, _ := repository.FindByVersion("apache", "1.0.0")

	assert.Equal(t, resources.Apache(), retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
	db.Exec("TRUNCATE TABLE latest_security_resources")
}
