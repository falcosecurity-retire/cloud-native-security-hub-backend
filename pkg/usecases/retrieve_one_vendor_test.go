package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryVendorRepository() vendor.Repository {
	return vendor.NewMemoryRepository(
		[]*vendor.Vendor{
			vendors.Apache(), vendors.Mongo(),
		},
	)
}

func TestReturnsOneVendorByID(t *testing.T) {
	useCase := usecases.RetrieveOneVendor{
		VendorRepository: memoryVendorRepository(),
		VendorID:         "apache",
	}

	result, _ := useCase.Execute()

	assert.Equal(t, vendors.Apache(), result)
}

func TestReturnsOneVendorNotFound(t *testing.T) {
	useCase := usecases.RetrieveOneVendor{
		VendorRepository: memoryVendorRepository(),
		VendorID:         "non-existent",
	}

	_, err := useCase.Execute()

	assert.Equal(t, vendor.ErrVendorNotFound, err)
}
