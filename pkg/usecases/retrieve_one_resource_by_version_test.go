package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepositoryByVersion() resource.Repository {
	apache := resources.Apache()
	apache.Version = "1.0.1"

	return resource.NewMemoryRepository(
		[]*resource.Resource{
			resources.Apache(),
			apache,
		},
	)
}

func TestReturnsOneResourceByVersion(t *testing.T) {
	useCase := usecases.RetrieveOneResourceByVersion{
		ResourceRepository: memoryResourceRepositoryByVersion(),
		ResourceID:         "apache",
		Version:            "1.0.1",
	}

	result, _ := useCase.Execute()

	expected := resources.Apache()
	expected.Version = "1.0.1"
	assert.Equal(t, expected, result)
}

func TestReturnsResourceByVersionNotFound(t *testing.T) {
	useCase := usecases.RetrieveOneResourceByVersion{
		ResourceRepository: memoryResourceRepositoryByVersion(),
		ResourceID:         "apache",
		Version:            "2.0.0",
	}

	_, err := useCase.Execute()

	assert.Equal(t, resource.ErrResourceNotFound, err)
}
