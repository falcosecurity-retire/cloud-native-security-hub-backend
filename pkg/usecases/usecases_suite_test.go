package usecases_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

func TestUsecases(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecases Suite")
}

func NewResourceRepository() resource.Repository {
	return resource.NewMemoryRepository(
		[]*resource.Resource{resources.Apache(), resources.MongoDB()},
	)
}

func NewVendorRepository() vendor.Repository {
	return vendor.NewMemoryRepository(
		[]*vendor.Vendor{vendors.Apache(), vendors.Mongo()},
	)
}
