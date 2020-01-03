package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveOneResource use case", func() {
	var useCase usecases.RetrieveOneResource

	BeforeEach(func() {
		useCase = usecases.RetrieveOneResource{
			ResourceRepository: NewResourceRepository(),
		}
	})

	It("returns one resource", func() {
		result, _ := useCase.Execute("apache", resource.FalcoRules)

		Expect(result).To(Equal(resources.Apache()))
	})

	Context("when resource does not exist", func() {
		It("returns resource not found error", func() {
			retrieved, err := useCase.Execute("notFound", resource.FalcoRules)

			Expect(retrieved).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})
