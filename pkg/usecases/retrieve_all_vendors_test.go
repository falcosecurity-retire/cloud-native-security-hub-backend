package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyVendorsRepository struct{}

func (resources *dummyVendorsRepository) FindAll() ([]*vendor.Resource, error) {
	return []*vendor.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	}, nil
}
func (resources *dummyVendorsRepository) FindById(id string) (*vendor.Resource, error) {
	return &vendor.Resource{
		Name: "Apache",
	}, nil
}

func TestReturnsAllVendors(t *testing.T) {
	useCase := RetrieveAllVendors{
		VendorRepository: &dummyVendorsRepository{},
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, resources, []*vendor.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	})
}
