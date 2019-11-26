package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
)

var _ = Describe("RetrieveAllResourcesFromVendor use case", func() {
	It("returns all the avaliable resources for a vendor", func() {
		useCase := usecases.RetrieveAllResourcesFromVendor{
			VendorID:           "apache",
			ResourceRepository: newResourceRepositoryWithoutMongoDB(),
			VendorRepository:   NewVendorRepository(),
		}

		retrieved, _ := useCase.Execute()

		Expect(retrieved).To(Equal([]*resource.Resource{resources.Apache()}))
	})

	Context("when vendor does not exist", func() {
		It("returns vendor not found error", func() {
			useCase := usecases.RetrieveAllResourcesFromVendor{
				VendorID:           "not-found",
				ResourceRepository: newResourceRepositoryWithoutMongoDB(),
				VendorRepository:   NewVendorRepository(),
			}

			retrieved, err := useCase.Execute()

			Expect(retrieved).To(BeEmpty())
			Expect(err).To(HaveOccurred())
		})
	})

	PContext("when vendor doesn't have resources", func() {
		It("returns an empty resource collection", func() {
			useCase := usecases.RetrieveAllResourcesFromVendor{
				VendorID:           "mongo",
				ResourceRepository: newResourceRepositoryWithoutMongoDB(),
				VendorRepository:   NewVendorRepository(),
			}

			retrieved, err := useCase.Execute()

			Expect(retrieved).To(BeEmpty())
			Expect(err).To(Succeed())
		})
	})
})

func newResourceRepositoryWithoutMongoDB() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.Apache()},
	)
}
