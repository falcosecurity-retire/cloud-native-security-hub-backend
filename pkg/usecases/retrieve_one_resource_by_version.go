package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResourceByVersion struct {
	ResourceRepository resource.Repository
}

func (r *RetrieveOneResourceByVersion) Execute(resourceID, kind, version string) (res *resource.Resource, err error) {
	return r.ResourceRepository.FindByVersion(
		resource.NewResourceID(resourceID, kind),
		version)
}
