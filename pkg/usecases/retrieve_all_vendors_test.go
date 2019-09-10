package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyVendorsRepository struct{}

func (resources *dummyVendorsRepository) FindAll() ([]*vendor.Vendor, error) {
	return []*vendor.Vendor{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	}, nil
}
func (resources *dummyVendorsRepository) FindById(id string) (*vendor.Vendor, error) {
	return &vendor.Vendor{
		Name: "Apache",
	}, nil
}

func TestReturnsAllVendors(t *testing.T) {
	useCase := RetrieveAllVendors{
		VendorRepository: &dummyVendorsRepository{},
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, resources, []*vendor.Vendor{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	})
}
