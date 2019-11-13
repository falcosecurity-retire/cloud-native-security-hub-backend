package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsAllVendors(t *testing.T) {
	existingVendors := []*vendor.Vendor{vendors.Apache(), vendors.Mongo()}
	vendorRepository := vendor.NewMemoryRepository(existingVendors)
	useCase := usecases.RetrieveAllVendors{VendorRepository: vendorRepository}

	retrieved, _ := useCase.Execute()

	assert.Equal(t, existingVendors, retrieved)
}
