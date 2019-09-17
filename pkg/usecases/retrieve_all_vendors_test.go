package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsAllVendors(t *testing.T) {
	vendorRepository := vendor.NewMemoryRepository(
		[]*vendor.Vendor{
			&vendor.Vendor{Name: "Apache"},
			&vendor.Vendor{Name: "Nginx"},
		},
	)
	useCase := RetrieveAllVendors{VendorRepository: vendorRepository}

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
