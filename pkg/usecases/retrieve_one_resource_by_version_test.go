package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveOneResourceByVersion use case", func() {
	var useCase usecases.RetrieveOneResourceByVersion

	BeforeEach(func() {
		useCase = usecases.RetrieveOneResourceByVersion{
			ResourceRepository: newResourceRepositoryWithVersions(),
		}
	})

	It("returns one resource", func() {
		result, _ := useCase.Execute("apache", resource.FalcoRules, "1.0.1")

		apacheWithSpecificVersion := resources.Apache()
		apacheWithSpecificVersion.Version = "1.0.1"
		Expect(result).To(Equal(apacheWithSpecificVersion))
	})

	Context("when version does not exist", func() {
		It("returns an error", func() {
			result, err := useCase.Execute("apache", resource.FalcoRules, "2.0.0")

			Expect(result).To(BeNil())
			Expect(err).To(MatchError(resource.ErrResourceNotFound))
		})
	})
})

func newResourceRepositoryWithVersions() resource.Repository {
	apache := resources.Apache()
	apache.Version = "1.0.1"

	return resource.NewMemoryRepository(
		[]*resource.Resource{
			resources.Apache(),
			apache,
		},
	)
}
