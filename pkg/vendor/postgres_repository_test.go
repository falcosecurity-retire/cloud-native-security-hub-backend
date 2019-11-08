package vendor

import (
	"database/sql"
	"os"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAVendor(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := NewPostgresRepository(db)

	vendor := apacheVendor()
	repository.Save(vendor)

	retrieved, _ := repository.FindById("apache")
	assert.Equal(t, vendor, retrieved)

	db.Exec("TRUNCATE TABLE vendors")
}

func apacheVendor() *Vendor {
	return &Vendor{
		ID:          "apache",
		Kind:        "Vendor",
		Name:        "Apache",
		Description: "# Apache Software Foundation\n",
		Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_Software_Foundation_Logo_%282016%29.svg/2560px-Apache_Software_Foundation_Logo_%282016%29.svg.png",
		Website:     "https://apache.org/",
	}
}

func mongodbVendor() *Vendor {
	return &Vendor{
		ID:          "mongo",
		Kind:        "Vendor",
		Name:        "Mongo",
		Description: "# MongoDB Inc.\n",
		Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/640px-MongoDB-Logo.svg.png",
		Website:     "https://mongodb.com/",
	}
}

func TestFindAllVendors(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := NewPostgresRepository(db)

	repository.Save(apacheVendor())
	repository.Save(mongodbVendor())

	retrieved, _ := repository.FindAll()
	assert.Equal(t, []*Vendor{apacheVendor(), mongodbVendor()}, retrieved)

	db.Exec("TRUNCATE TABLE vendors")
}
