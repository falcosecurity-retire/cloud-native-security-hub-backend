package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneVendor struct {
	VendorID         string
	VendorRepository resource.Repository
}

func (useCase *RetrieveOneVendor) Execute() (res *resource.Resource, err error) {
	return useCase.VendorRepository.FindById(useCase.VendorID)
}
