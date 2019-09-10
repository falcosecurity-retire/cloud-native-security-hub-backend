package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepositoryFromVendor() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{
			{
				Name:   "Falco profile for Nginx",
				Vendor: "Nginx",
			},
			{
				Name:   "Grafana Dashboard for Traefik",
				Vendor: "Traefik",
			},
		},
	)
}

func memoryVendorRepositoryFromVendor() vendor.Repository {
	return vendor.NewMemoryRepository(
		[]*vendor.Vendor{
			{
				ID:   "apache",
				Name: "Apache",
			},
			{
				ID:   "nginx",
				Name: "Nginx",
			},
			{
				ID:   "traefik",
				Name: "Traefik",
			},
		},
	)
}

func TestReturnsAllResourcesFromVendor(t *testing.T) {
	useCase := RetrieveAllResourcesFromVendor{
		VendorID:           "Nginx",
		ResourceRepository: memoryResourceRepositoryFromVendor(),
		VendorRepository:   memoryVendorRepositoryFromVendor(),
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, []*resource.Resource{
		{
			Name:   "Falco profile for Nginx",
			Vendor: "Nginx",
		},
	}, resources)
}

func TestReturnsVendorNotFoundResourcesFromVendor(t *testing.T) {
	useCase := RetrieveAllResourcesFromVendor{
		VendorID:           "not-found",
		ResourceRepository: memoryResourceRepositoryFromVendor(),
		VendorRepository:   memoryVendorRepositoryFromVendor(),
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}

func TestReturnsResourcesNotFoundResourcesFromVendor(t *testing.T) {
	useCase := RetrieveAllResourcesFromVendor{
		VendorID:           "apache",
		ResourceRepository: memoryResourceRepositoryFromVendor(),
		VendorRepository:   memoryVendorRepositoryFromVendor(),
	}

	_, err := useCase.Execute()

	assert.Error(t, err) //vendor exists but has no resources
}
