package usecases

import (
	"fmt"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type dummyResourcesRepositoryForOne struct{}

func (resources *dummyResourcesRepositoryForOne) FindAll() ([]*resource.Resource, error) {
	return []*resource.Resource{
		{
			Kind:   resource.FALCO_RULE,
			Name:   "Falco profile for Nginx",
			Vendor: "Nginx",
		},
		{
			Kind:   "GrafanaDashboard",
			Name:   "Grafana Dashboard for Traefik",
			Vendor: "Traefik",
		},
	}, nil
}

func (resources *dummyResourcesRepositoryForOne) FindById(id string) (*resource.Resource, error) {
	all, err := resources.FindAll()
	if err != nil {
		return nil, err
	}
	idToFind := strings.ToLower(id)
	for _, res := range all {
		resName := strings.ToLower(res.Name)
		resHash := strings.ToLower(res.Hash())
		if resName == idToFind || resHash == idToFind {
			return res, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func TestReturnsOneResource(t *testing.T) {
	useCase := RetrieveOneResource{
		ResourceRepository: &dummyResourcesRepositoryForOne{},
		Hash:               "bekiisotdwhvmetchrwp",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, &resource.Resource{
		Kind:   resource.FALCO_RULE,
		Name:   "Falco profile for Nginx",
		Vendor: "Nginx",
	}, res)
}

func TestReturnsResourceNotFound(t *testing.T) {
	useCase := RetrieveOneResource{
		ResourceRepository: &dummyResourcesRepositoryForOne{},
		Hash:               "notFound",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
