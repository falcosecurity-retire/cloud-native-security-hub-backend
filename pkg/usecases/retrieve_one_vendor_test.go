package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

var _ = Describe("RetrieveOneVendor use case", func() {
	var useCase usecases.RetrieveOneVendor

	BeforeEach(func() {
		useCase = usecases.RetrieveOneVendor{
			VendorRepository: NewVendorRepository(),
		}
	})

	It("returns one vendor", func() {
		result, _ := useCase.Execute("apache")

		Expect(result).To(Equal(vendors.Apache()))
	})

	Context("when vendor does not exist", func() {
		It("returns vendor not found error", func() {
			retrieved, err := useCase.Execute("nonExistent")

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(vendor.ErrVendorNotFound))
		})
	})
})
