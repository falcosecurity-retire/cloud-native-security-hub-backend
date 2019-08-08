package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyResourcesRepository struct{}

func (resources *dummyResourcesRepository) All() ([]resource.Resource, error) {
	return []resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	}, nil
}

func TestReturnsAllResources(t *testing.T) {
	useCase := RetrieveAllResources{
		ResourceRepository: &dummyResourcesRepository{},
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, resources, []resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	})
}
