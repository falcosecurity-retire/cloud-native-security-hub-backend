package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyVendorRepository struct{}

func (resources *dummyVendorRepository) All() ([]resource.Resource, error) {
	return []resource.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	}, nil
}

func TestReturnsOneVendor(t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: &dummyVendorRepository{},
		VendorID:         "apache",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, res, resource.Resource{
		Name: "Apache",
	})
}

func TestReturnsOneVendorNotFound (t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: &dummyVendorRepository{},
		VendorID:         "non-existent",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}