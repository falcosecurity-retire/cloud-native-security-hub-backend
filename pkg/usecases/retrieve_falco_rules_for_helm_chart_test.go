package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepositoryWithRules() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.Apache(), resources.MongoDB()},
	)
}

func TestReturnsFalcoRulesForHelmChart(t *testing.T) {
	useCase := usecases.RetrieveFalcoRulesForHelmChart{
		ResourceRepository: memoryResourceRepositoryWithRules(),
		ResourceID:         "apache",
	}

	result, _ := useCase.Execute()

	expected := `customRules:
  rules-apache.yaml: |
    - macro: apache_consider_syscalls
      condition: (evt.num < 0)
`
	assert.Equal(t, expected, string(result))
}

func TestFalcoRulesForHelmChartReturnsNotFound(t *testing.T) {
	useCase := usecases.RetrieveFalcoRulesForHelmChart{
		ResourceRepository: memoryResourceRepositoryWithRules(),
		ResourceID:         "notFound",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
