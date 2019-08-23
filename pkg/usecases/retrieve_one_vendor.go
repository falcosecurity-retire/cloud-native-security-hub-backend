package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

type RetrieveOneVendor struct {
	VendorID         string
	VendorRepository vendor.Repository
}

func (useCase *RetrieveOneVendor) Execute() (res *vendor.Resource, err error) {
	return useCase.VendorRepository.FindById(useCase.VendorID)
}
