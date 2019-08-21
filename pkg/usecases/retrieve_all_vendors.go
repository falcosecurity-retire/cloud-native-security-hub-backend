package usecases

import "cloud-native-visibility-hub/pkg/resource"

type RetrieveAllVendors struct {
	VendorRepository resource.Repository
}

func (useCase *RetrieveAllVendors) Execute() ([]resource.Resource, error) {
	return useCase.VendorRepository.All()
}
