package usecases

import (
	"fmt"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type dummyResourcesRepositoryForOneRaw struct{}

func (resources *dummyResourcesRepositoryForOneRaw) FindAll() ([]*resource.Resource, error) {
	return []*resource.Resource{
		{
			Kind:   resource.FALCO_RULE,
			Name:   "Falco profile for Nginx",
			Vendor: "Nginx",
			Rules: []*resource.FalcoRuleData{
				{Raw: "nginxRule"},
			},
		},
		{
			Kind:   "GrafanaDashboard",
			Name:   "Grafana Dashboard for Traefik",
			Vendor: "Traefik",
			Rules: []*resource.FalcoRuleData{
				{Raw: "traefikRule"},
			},
		},
	}, nil
}

func (resources *dummyResourcesRepositoryForOneRaw) FindById(id string) (*resource.Resource, error) {
	all, err := resources.FindAll()
	if err != nil {
		return nil, err
	}
	idToFind := strings.ToLower(id)
	for _, res := range all {
		resourceName := strings.ToLower(res.Name)
		resourceID := strings.ToLower(res.ID)
		if resourceName == idToFind || resourceID == idToFind {
			return res, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func TestReturnsOneRawResource(t *testing.T) {
	useCase := RetrieveOneRawResource{
		ResourceRepository: &dummyResourcesRepositoryForOneRaw{},
		ResourceID:         "Falco profile for Nginx",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, []byte("nginxRule"), res.Raw())
}

func TestReturnsResourceRawNotFound(t *testing.T) {
	useCase := RetrieveOneRawResource{
		ResourceRepository: &dummyResourcesRepositoryForOneRaw{},
		ResourceID:         "notFound",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
