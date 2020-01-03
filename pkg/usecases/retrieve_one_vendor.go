package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

type RetrieveOneVendor struct {
	VendorRepository vendor.Repository
}

func (r *RetrieveOneVendor) Execute(vendorID string) (*vendor.Vendor, error) {
	return r.VendorRepository.FindById(vendorID)
}
