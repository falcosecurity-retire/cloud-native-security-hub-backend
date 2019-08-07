package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
)

type RetrieveAllResources struct {
	Resources resource.ResourceRepository
}

func (useCase *RetrieveAllResources) Execute() ([]resource.Resource, error) {
	return useCase.Resources.All()
}
