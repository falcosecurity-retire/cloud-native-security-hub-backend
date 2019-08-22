package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"fmt"
	"strings"
)

type RetrieveOneVendor struct {
	VendorID         string
	VendorRepository resource.Repository
}

func (useCase *RetrieveOneVendor) Execute() (res *resource.Resource, err error) {
	vendors, err := useCase.VendorRepository.FindAll()
	if err != nil {
		return
	}

	for _, vendor := range vendors {
		vendorName := strings.ToLower(vendor.Name)
		vendorID := strings.ToLower(vendor.Hash())
		vendorToLookFor := strings.ToLower(useCase.VendorID)
		if vendorName == vendorToLookFor || vendorID == vendorToLookFor {
			res = vendor
			return
		}
	}

	err = fmt.Errorf("not found")
	return
}
