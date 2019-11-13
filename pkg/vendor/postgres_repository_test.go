package vendor_test

import (
	"database/sql"
	"os"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAVendor(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := vendor.NewPostgresRepository(db)

	repository.Save(vendors.Apache())

	retrieved, _ := repository.FindById("apache")
	assert.Equal(t, vendors.Apache(), retrieved)

	db.Exec("TRUNCATE TABLE vendors")
}

func TestFindAllVendors(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := vendor.NewPostgresRepository(db)

	repository.Save(vendors.Apache())
	repository.Save(vendors.Mongo())

	retrieved, _ := repository.FindAll()

	assert.Equal(t, []*vendor.Vendor{vendors.Apache(), vendors.Mongo()}, retrieved)

	db.Exec("TRUNCATE TABLE vendors")
}

func TestFindVendorByIdDoesntFindTheVendor(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := vendor.NewPostgresRepository(db)

	retrieved, err := repository.FindById("non existent id")

	assert.Nil(t, retrieved)
	assert.Equal(t, vendor.ErrVendorNotFound, err)
}
