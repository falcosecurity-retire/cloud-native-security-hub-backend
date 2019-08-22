package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveAllResources struct {
	ResourceRepository resource.Repository
}

func (useCase *RetrieveAllResources) Execute() ([]*resource.Resource, error) {
	return useCase.ResourceRepository.FindAll()
}
