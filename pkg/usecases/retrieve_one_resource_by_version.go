package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResourceByVersion struct {
	ResourceRepository resource.Repository
	ResourceID         string
	Version            string
}

func (useCase *RetrieveOneResourceByVersion) Execute() (res *resource.Resource, err error) {
	return useCase.ResourceRepository.FindByVersion(useCase.ResourceID, useCase.Version)
}
