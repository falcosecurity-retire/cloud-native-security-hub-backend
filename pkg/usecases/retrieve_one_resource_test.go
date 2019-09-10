package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepository() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{
			&resource.Resource{
				Kind:   resource.FALCO_RULE,
				Name:   "Falco profile for Nginx",
				Vendor: "Nginx",
				ID:     "nginx",
			},
			&resource.Resource{
				Kind:   resource.FALCO_RULE,
				Name:   "Falco profile for Traefik",
				Vendor: "Traefik",
				ID:     "traefik",
			},
		},
	)
}

func TestReturnsOneResource(t *testing.T) {
	useCase := RetrieveOneResource{
		ResourceRepository: memoryResourceRepository(),
		ResourceID:         "nginx",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, &resource.Resource{
		Kind:   resource.FALCO_RULE,
		Name:   "Falco profile for Nginx",
		Vendor: "Nginx",
		ID:     "nginx",
	}, res)
}

func TestReturnsResourceNotFound(t *testing.T) {
	useCase := RetrieveOneResource{
		ResourceRepository: memoryResourceRepository(),
		ResourceID:         "notFound",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
