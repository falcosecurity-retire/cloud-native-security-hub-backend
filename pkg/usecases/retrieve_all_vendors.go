package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveAllVendors struct {
	VendorRepository resource.Repository
}

func (useCase *RetrieveAllVendors) Execute() ([]resource.Resource, error) {
	return useCase.VendorRepository.All()
}
