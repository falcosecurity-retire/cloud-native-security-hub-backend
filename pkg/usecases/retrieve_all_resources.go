package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
)

type RetrieveAllResources struct {
	ResourceRepository resource.Repository
}

func (useCase *RetrieveAllResources) Execute() ([]resource.Resource, error) {
	return useCase.ResourceRepository.All()
}
