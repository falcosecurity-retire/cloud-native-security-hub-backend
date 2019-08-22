package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyResourcesRepository struct{}

func (resources *dummyResourcesRepository) FindAll() ([]*resource.Resource, error) {
	return []*resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	}, nil
}

func (resources *dummyResourcesRepository) FindById(id string) (*resource.Resource, error) {
	return &resource.Resource{

		Name: "Falco profile for Nginx",
	}, nil
}

func TestReturnsAllResources(t *testing.T) {
	useCase := RetrieveAllResources{
		ResourceRepository: &dummyResourcesRepository{},
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, []*resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	}, resources)
}
