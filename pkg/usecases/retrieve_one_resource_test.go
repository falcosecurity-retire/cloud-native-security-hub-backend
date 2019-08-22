package usecases;

import (
	"cloud-native-visibility-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyResourcesRepositoryForOne struct{}

func (resources *dummyResourcesRepositoryForOne) All() ([]resource.Resource, error) {
	return []resource.Resource{
		{
			Kind:       resource.FALCO_RULE,
			Name:       "Falco profile for Nginx",
			Vendor:     "Nginx",
		},
		{
			Kind:       "GrafanaDashboard",
			Name:       "Grafana Dashboard for Traefik",
			Vendor:     "Traefik",
		},
	}, nil
}

func TestReturnsOneResource(t *testing.T) {
	useCase := RetrieveOneResource{
		ResourceRepository: &dummyResourcesRepositoryForOne{},
		Hash:               "bekiisotdwhvmetchrwp",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, res, resource.Resource{
		Kind:       resource.FALCO_RULE,
		Name:       "Falco profile for Nginx",
		Vendor:     "Nginx",
	})
}

func TestReturnsResourceNotFound(t *testing.T) {
	useCase := RetrieveOneResource{
		ResourceRepository: &dummyResourcesRepositoryForOne{},
		Hash:               "notFound",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
