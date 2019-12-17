package usecases_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
)

var _ = Describe("RetrieveFalcoRulesForHelmChartByVersion use case", func() {
	It("returns the rules for being used with the Helm chart", func() {
		useCase := usecases.RetrieveFalcoRulesForHelmChartByVersion{
			ResourceRepository: NewResourceRepository(),
			ResourceID:         "apache",
			Version:            "1.0.0",
		}

		result, _ := useCase.Execute()

		expected := `customRules:
  rules-apache.yaml: |
    - macro: apache_consider_syscalls
      condition: (evt.num < 0)
`
		Expect(expected).To(Equal(string(result)))
	})

	Context("when resource doesn't exist", func() {
		It("it returns a resource not found error", func() {
			useCase := usecases.RetrieveFalcoRulesForHelmChartByVersion{
				ResourceRepository: NewResourceRepository(),
				ResourceID:         "notFound",
			}

			_, err := useCase.Execute()

			Expect(err).To(HaveOccurred())
		})
	})

	Context("when version doesn't exist", func() {
		It("it returns a resource not found error", func() {
			useCase := usecases.RetrieveFalcoRulesForHelmChartByVersion{
				ResourceRepository: NewResourceRepository(),
				ResourceID:         "apache",
				Version:            "3.0.0",
			}

			_, err := useCase.Execute()

			Expect(err).To(HaveOccurred())
		})
	})
})
