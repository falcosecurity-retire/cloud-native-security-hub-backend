package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

var _ = Describe("RetrieveAllVendors use case", func() {
	It("returns all the available vendors", func() {
		existingVendors := []*vendor.Vendor{vendors.Apache(), vendors.Mongo()}
		useCase := usecases.RetrieveAllVendors{VendorRepository: NewVendorRepository()}

		retrieved, _ := useCase.Execute()

		Expect(retrieved).To(Equal(existingVendors))
	})
})
