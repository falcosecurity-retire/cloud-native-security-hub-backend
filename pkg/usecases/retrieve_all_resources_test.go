package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveAllResources use case", func() {
	It("returns all the available resources", func() {
		existingResources := []*resource.Resource{resources.Apache(), resources.MongoDB()}
		useCase := usecases.RetrieveAllResources{ResourceRepository: NewResourceRepository()}

		retrieved, _ := useCase.Execute()

		Expect(retrieved).To(Equal(existingResources))
	})
})
