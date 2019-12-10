package resource_test

import (
	"database/sql"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Postgres Updater", func() {
	var (
		repository resource.Repository
		updater    resource.Updater
	)

	BeforeEach(func() {
		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		repository = resource.NewPostgresRepository(db)
		updater = resource.NewPostgresUpdater(db)

		db.Exec("TRUNCATE TABLE security_resources")
		db.Exec("TRUNCATE TABLE latest_security_resources")

		repository.Save(resources.Apache())
	})

	It("increments download count for Apache", func() {
		first, _ := repository.FindById(resources.Apache().ID)

		updater.IncrementDownloadCountFor(resources.Apache().ID)

		second, _ := repository.FindById(resources.Apache().ID)

		Expect(second.DownloadCount).To(Equal(first.DownloadCount + 1))
	})
})
