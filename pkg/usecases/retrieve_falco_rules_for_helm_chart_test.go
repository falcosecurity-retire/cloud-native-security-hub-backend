package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepositoryWithRules() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{
			{
				ID:     "nginx",
				Kind:   resource.FALCO_RULE,
				Name:   "Falco profile for Nginx",
				Vendor: "Nginx",
				Rules: []*resource.FalcoRuleData{
					{Raw: "nginxRule"},
				},
			},
			{
				ID:     "traefik",
				Kind:   "GrafanaDashboard",
				Name:   "Grafana Dashboard for Traefik",
				Vendor: "Traefik",
				Rules: []*resource.FalcoRuleData{
					{Raw: "traefikRule"},
				},
			},
		},
	)
}

func TestReturnsFalcoRulesForHelmChart(t *testing.T) {
	useCase := RetrieveFalcoRulesForHelmChart{
		ResourceRepository: memoryResourceRepositoryWithRules(),
		ResourceID:         "nginx",
	}

	result, _ := useCase.Execute()
	expected := `customRules:
  rules-nginx.yaml: nginxRule
`

	assert.Equal(t, expected, string(result))
}

func TestFalcoRulesForHelmChartReturnsNotFound(t *testing.T) {
	useCase := RetrieveFalcoRulesForHelmChart{
		ResourceRepository: memoryResourceRepositoryWithRules(),
		ResourceID:         "notFound",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
