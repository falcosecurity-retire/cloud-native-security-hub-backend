package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveOneResource use case", func() {
	It("returns one resource", func() {
		useCase := usecases.RetrieveOneResource{
			ResourceRepository: NewResourceRepository(),
			ResourceID:         "apache",
			Kind:               resource.FalcoRules,
		}

		result, _ := useCase.Execute()

		Expect(result).To(Equal(resources.Apache()))
	})

	Context("when resource does not exist", func() {
		It("returns resource not found error", func() {
			useCase := usecases.RetrieveOneResource{
				ResourceRepository: NewResourceRepository(),
				ResourceID:         "notFound",
				Kind:               resource.FalcoRules,
			}

			retrieved, err := useCase.Execute()

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})
