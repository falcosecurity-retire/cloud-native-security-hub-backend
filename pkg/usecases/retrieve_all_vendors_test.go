package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyVendorsRepository struct{}

func (resources *dummyVendorsRepository) FindAll() ([]*resource.Resource, error) {
	return []*resource.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	}, nil
}
func (resources *dummyVendorsRepository) FindById(id string) (*resource.Resource, error) {
	return &resource.Resource{
		Name: "Apache",
	}, nil
}

func TestReturnsAllVendors(t *testing.T) {
	useCase := RetrieveAllVendors{
		VendorRepository: &dummyVendorsRepository{},
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, resources, []*resource.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	})
}
