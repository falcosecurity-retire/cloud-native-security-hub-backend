package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyVendorRepository struct{}

func (resources *dummyVendorRepository) FindAll() ([]*resource.Resource, error) {
	return []*resource.Resource{
		{
			Name: "Apache",
		},
		{
			Name: "Nginx",
		},
	}, nil
}

func (resources *dummyVendorRepository) FindById(id string) (*resource.Resource, error) {
	return &resource.Resource{
		Name: "Apache",
	}, nil
}

func TestReturnsOneVendorByName(t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: &dummyVendorRepository{},
		VendorID:         "apache",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, &resource.Resource{
		Name: "Apache",
	}, res)
}

func TestReturnsOneVendorByHash(t *testing.T) {
	repository := &dummyVendorRepository{}
	all, _ := repository.FindAll()
	expected := all[0]

	useCase := RetrieveOneVendor{
		VendorRepository: repository,
		VendorID:         expected.Hash(),
	}

	res, _ := useCase.Execute()

	assert.Equal(t, expected, res)
}

func TestReturnsOneVendorNotFound(t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: &dummyVendorRepository{},
		VendorID:         "non-existent",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
