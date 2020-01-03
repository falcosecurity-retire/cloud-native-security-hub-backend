package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveFalcoRulesForHelmChartByVersion use case", func() {
	var useCase usecases.RetrieveFalcoRulesForHelmChartByVersion

	BeforeEach(func() {
		useCase = usecases.RetrieveFalcoRulesForHelmChartByVersion{
			ResourceRepository: NewResourceRepository(),
		}
	})

	It("returns the rules for being used with the Helm chart", func() {
		result, _ := useCase.Execute("apache", "1.0.0")

		expected := `customRules:
  rules-apache.yaml: |
    - macro: apache_consider_syscalls
      condition: (evt.num < 0)
`
		Expect(expected).To(Equal(string(result)))
	})

	Context("when resource doesn't exist", func() {
		It("it returns a resource not found error", func() {
			_, err := useCase.Execute("notFound", "1.0.0")

			Expect(err).To(HaveOccurred())
		})
	})

	Context("when version doesn't exist", func() {
		It("it returns a resource not found error", func() {
			_, err := useCase.Execute("apache", "3.0.0")

			Expect(err).To(HaveOccurred())
		})
	})
})
