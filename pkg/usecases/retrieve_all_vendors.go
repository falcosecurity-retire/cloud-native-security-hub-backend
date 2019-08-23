package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"

type RetrieveAllVendors struct {
	VendorRepository vendor.Repository
}

func (useCase *RetrieveAllVendors) Execute() ([]*vendor.Resource, error) {
	return useCase.VendorRepository.FindAll()
}
