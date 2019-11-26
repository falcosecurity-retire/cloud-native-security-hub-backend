package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

var _ = Describe("RetrieveOneVendor use case", func() {
	It("returns one vendor", func() {
		useCase := usecases.RetrieveOneVendor{
			VendorRepository: NewVendorRepository(),
			VendorID:         "apache",
		}

		result, _ := useCase.Execute()

		Expect(result).To(Equal(vendors.Apache()))
	})

	Context("when vendor does not exist", func() {
		It("returns vendor not found error", func() {
			useCase := usecases.RetrieveOneVendor{
				VendorRepository: NewVendorRepository(),
				VendorID:         "non-existent",
			}

			retrieved, err := useCase.Execute()

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(vendor.ErrVendorNotFound))
		})
	})
})
