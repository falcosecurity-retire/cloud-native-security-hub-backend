package vendor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"

	"database/sql"
	"os"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

var _ = Describe("Postgres Vendor Repository", func() {
	var repository vendor.Repository

	BeforeEach(func() {
		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		repository = vendor.NewPostgresRepository(db)

		db.Exec("TRUNCATE TABLE vendors")
	})

	It("saves a new vendor", func() {
		repository.Save(vendors.Apache())

		retrieved, _ := repository.FindById("apache")

		Expect(retrieved).To(Equal(vendors.Apache()))
	})

	It("retrieves all existent vendors", func() {
		repository.Save(vendors.Apache())
		repository.Save(vendors.Mongo())

		retrieved, _ := repository.FindAll()

		Expect(retrieved).To(Equal([]*vendor.Vendor{vendors.Apache(), vendors.Mongo()}))
	})

	Context("when querying by id", func() {
		Context("and vendor does not exist", func() {
			It("returns an error", func() {
				retrieved, err := repository.FindById("non existent id")

				Expect(retrieved).To(BeNil())
				Expect(err).To(MatchError(vendor.ErrVendorNotFound))
			})
		})
	})
})
