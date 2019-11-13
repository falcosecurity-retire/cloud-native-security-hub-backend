package usecases

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

func memoryResourceRepositoryByVersion() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{
			&resource.Resource{
				Kind:    resource.FALCO_RULE,
				Name:    "Falco profile for Nginx",
				Vendor:  "Nginx",
				ID:      "nginx",
				Version: "1.0.0",
			},
			&resource.Resource{
				Kind:    resource.FALCO_RULE,
				Name:    "Falco profile for Nginx",
				Vendor:  "Nginx",
				ID:      "nginx",
				Version: "1.0.1",
			},
		},
	)
}

func TestReturnsOneResourceByVersion(t *testing.T) {
	useCase := RetrieveOneResourceByVersion{
		ResourceRepository: memoryResourceRepositoryByVersion(),
		ResourceID:         "nginx",
		Version:            "1.0.1",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, &resource.Resource{
		Kind:    resource.FALCO_RULE,
		Name:    "Falco profile for Nginx",
		Vendor:  "Nginx",
		ID:      "nginx",
		Version: "1.0.1",
	}, res)
}

func TestReturnsResourceByVersionNotFound(t *testing.T) {
	useCase := RetrieveOneResourceByVersion{
		ResourceRepository: memoryResourceRepositoryByVersion(),
		ResourceID:         "apache",
		Version:            "2.0.0",
	}

	_, err := useCase.Execute()

	assert.Equal(t, resource.ErrResourceNotFound, err)
}
