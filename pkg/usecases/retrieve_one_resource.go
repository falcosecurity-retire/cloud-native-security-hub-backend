package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResource struct {
	ResourceRepository resource.Repository
}

func (r *RetrieveOneResource) Execute(resourceID, kind string) (res *resource.Resource, err error) {
	return r.ResourceRepository.FindById(resource.NewResourceID(resourceID, kind))
}
