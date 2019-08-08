package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
)

type RetrieveAllResources struct {
	Resources resource.Repository
}

func (useCase *RetrieveAllResources) Execute() ([]resource.Resource, error) {
	return useCase.Resources.All()
}
