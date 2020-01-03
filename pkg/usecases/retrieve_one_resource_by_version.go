package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResourceByVersion struct {
	ResourceRepository resource.Repository
	ResourceID         string
	Kind               string
	Version            string
}

func (useCase *RetrieveOneResourceByVersion) Execute() (res *resource.Resource, err error) {
	return useCase.ResourceRepository.FindByVersion(
		resource.NewResourceID(useCase.ResourceID, useCase.Kind),
		useCase.Version)
}
