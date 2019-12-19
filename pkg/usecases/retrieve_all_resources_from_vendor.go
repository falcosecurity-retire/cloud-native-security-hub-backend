package usecases

import (
	"fmt"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

type RetrieveAllResourcesFromVendor struct {
	VendorID           string
	VendorRepository   vendor.Repository
	ResourceRepository resource.Repository
}

func (useCase *RetrieveAllResourcesFromVendor) Execute() (res []*resource.Resource, err error) {
	vendor, err := useCase.VendorRepository.FindById(useCase.VendorID)
	if err != nil {
		return
	}

	resources, err := useCase.ResourceRepository.FindAll()
	if err != nil {
		return
	}

	for _, r := range resources {
		if vendor.Name == r.Vendor {
			res = append(res, r)
		}
	}

	if len(res) == 0 {
		err = fmt.Errorf("no resources available for this vendor")
	}

	return
}
