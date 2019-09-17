package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryVendorRepository() vendor.Repository {
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
		},
	)
}

func TestReturnsOneVendorByName(t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: memoryVendorRepository(),
		VendorID:         "apache",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, &vendor.Vendor{
		ID:   "apache",
		Name: "Apache",
	}, res)
}

func TestReturnsOneVendorByID(t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: memoryVendorRepository(),
		VendorID:         "apache",
	}

	res, _ := useCase.Execute()

	assert.Equal(t, &vendor.Vendor{
		ID:   "apache",
		Name: "Apache",
	}, res)
}

func TestReturnsOneVendorNotFound(t *testing.T) {
	useCase := RetrieveOneVendor{
		VendorRepository: memoryVendorRepository(),
		VendorID:         "non-existent",
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}
