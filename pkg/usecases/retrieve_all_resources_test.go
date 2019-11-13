package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsAllResources(t *testing.T) {
	existingResources := []*resource.Resource{resources.Apache(), resources.MongoDB()}
	resourceRepository := resource.NewMemoryRepository(existingResources)
	useCase := usecases.RetrieveAllResources{ResourceRepository: resourceRepository}

	retrieved, _ := useCase.Execute()

	assert.Equal(t, existingResources, retrieved)
}
