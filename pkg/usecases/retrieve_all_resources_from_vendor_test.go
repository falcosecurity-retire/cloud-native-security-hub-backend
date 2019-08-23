package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyResourcesFromVendor_Resources struct{}

func (resources *dummyResourcesFromVendor_Resources) FindAll() ([]*resource.Resource, error) {
	return []*resource.Resource{
		{
			Name:   "Falco profile for Nginx",
			Vendor: "Nginx",
		},
		{
			Name:   "Grafana Dashboard for Traefik",
			Vendor: "Traefik",
		},
	}, nil
}

func (resources *dummyResourcesFromVendor_Resources) FindById(id string) (*resource.Resource, error) {
	return &resource.Resource{
		Name: "Falco profile for Nginx",
	}, nil
}

type dummyResourcesFromVendor_Vendors struct{}

func (resources *dummyResourcesFromVendor_Vendors) FindAll() ([]*vendor.Resource, error) {
	return []*vendor.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
		{
			Name: "Traefik",
		},
	}, nil
}
func (resources *dummyResourcesFromVendor_Vendors) FindById(id string) (*vendor.Resource, error) {
	return &vendor.Resource{
		Name: id,
	}, nil
}

func TestReturnsAllResourcesFromVendor(t *testing.T) {
	useCase := RetrieveAllResourcesFromVendor{
		VendorID:           "Nginx",
		ResourceRepository: &dummyResourcesFromVendor_Resources{},
		VendorRepository:   &dummyResourcesFromVendor_Vendors{},
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
		ResourceRepository: &dummyResourcesFromVendor_Resources{},
		VendorRepository:   &dummyResourcesFromVendor_Vendors{},
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}

func TestReturnsResourcesNotFoundResourcesFromVendor(t *testing.T) {
	useCase := RetrieveAllResourcesFromVendor{
		VendorID:           "apache",
		ResourceRepository: &dummyResourcesFromVendor_Resources{},
		VendorRepository:   &dummyResourcesFromVendor_Vendors{},
	}

	_, err := useCase.Execute()

	assert.Error(t, err) //vendor exists but has no resources
}
