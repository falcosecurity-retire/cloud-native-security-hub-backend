package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsAllResources(t *testing.T) {
	resourceRepository := resource.MemoryRepository{}
	resourceRepository.Add(resource.Resource{Name: "Falco profile for Nginx"})
	resourceRepository.Add(resource.Resource{Name: "Falco profile for Grafana"})

	useCase := RetrieveAllResources{ResourceRepository: &resourceRepository}

	resources, _ := useCase.Execute()

	assert.Equal(t, []*resource.Resource{
		{Name: "Falco profile for Nginx"},
		{Name: "Falco profile for Grafana"},
	}, resources)
}
