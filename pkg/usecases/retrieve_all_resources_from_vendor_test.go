package usecases_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/usecases"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func memoryResourceRepositoryFromVendor() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.Apache(), resources.MongoDB()},
	)
}

func memoryVendorRepositoryFromVendor() vendor.Repository {
	return vendor.NewMemoryRepository(
		[]*vendor.Vendor{vendors.Apache(), vendors.Mongo()},
	)
}

func TestReturnsAllResourcesFromVendor(t *testing.T) {
	useCase := usecases.RetrieveAllResourcesFromVendor{
		VendorID:           "apache",
		ResourceRepository: memoryResourceRepositoryFromVendor(),
		VendorRepository:   memoryVendorRepositoryFromVendor(),
	}

	retrieved, _ := useCase.Execute()

	assert.Equal(t, []*resource.Resource{resources.Apache()}, retrieved)
}

func TestReturnsVendorNotFoundResourcesFromVendor(t *testing.T) {
	useCase := usecases.RetrieveAllResourcesFromVendor{
		VendorID:           "not-found",
		ResourceRepository: memoryResourceRepositoryFromVendor(),
		VendorRepository:   memoryVendorRepositoryFromVendor(),
	}

	_, err := useCase.Execute()

	assert.Error(t, err)
}

// FIXME: If the vendor exists and doesn't have resources it should return an empty array
// This is not an error
func TestReturnsResourcesNotFoundResourcesFromVendor(t *testing.T) {
	useCase := usecases.RetrieveAllResourcesFromVendor{
		VendorID:           "nginx",
		ResourceRepository: memoryResourceRepositoryFromVendor(),
		VendorRepository:   memoryVendorRepositoryFromVendor(),
	}

	_, err := useCase.Execute()

	assert.Error(t, err) //vendor exists but has no resources
}
