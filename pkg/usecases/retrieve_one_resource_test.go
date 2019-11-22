package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepository() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{
			resources.Apache(), resources.MongoDB(),
		},
	)
}

func TestReturnsOneResource(t *testing.T) {
	useCase := usecases.RetrieveOneResource{
		ResourceRepository: memoryResourceRepository(),
		ResourceID:         "apache",
	}

	result, _ := useCase.Execute()

	assert.Equal(t, resources.Apache(), result)
}

func TestReturnsResourceNotFound(t *testing.T) {
	useCase := usecases.RetrieveOneResource{
		ResourceRepository: memoryResourceRepository(),
		ResourceID:         "notFound",
	}

	_, err := useCase.Execute()

	assert.Equal(t, resource.ErrResourceNotFound, err)
}
