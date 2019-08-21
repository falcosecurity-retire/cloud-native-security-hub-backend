package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyVendorsRepository struct{}

func (resources *dummyVendorsRepository) All() ([]resource.Resource, error) {
	return []resource.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	}, nil
}

func TestReturnsAllVendors(t *testing.T) {
	useCase := RetrieveAllVendors{
		VendorRepository: &dummyVendorsRepository{},
	}

	resources, _ := useCase.Execute()

	assert.Equal(t, resources, []resource.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	})
}
