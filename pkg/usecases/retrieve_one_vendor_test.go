package usecases

import (
	"fmt"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"strings"
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
	all, err := resources.FindAll()
	if err != nil {
		return nil, err
	}
	idToFind := strings.ToLower(id)
	for _, res := range all {
		resName := strings.ToLower(res.Name)
		resHash := strings.ToLower(res.Hash())
		if resName == idToFind || resHash == idToFind {
			return res, nil
		}
	}
	return nil, fmt.Errorf("not found")
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
