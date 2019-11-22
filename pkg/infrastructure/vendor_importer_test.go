package infrastructure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/infrastructure"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

var _ = Describe("Vendor importation from YAML files", func() {
	It("walks a directory and extract resources", func() {
		path := "../../test/fixtures/vendors"
		parsed, _ := infrastructure.GetVendorsFromPath(path)

		Expect(parsed).To(Equal([]*vendor.Vendor{
			vendors.Apache(),
			vendors.Mongo(),
		}))
	})

	Context("when path doesn't exist", func() {
		It("returns an error", func() {
			nonExistentPath := "../foo"

			parsed, err := infrastructure.GetVendorsFromPath(nonExistentPath)

			Expect(parsed).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})
})
