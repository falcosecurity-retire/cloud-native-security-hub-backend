package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
)

type Factory interface {
	NewRetrieveAllResourcesUseCase() *RetrieveAllResources

	NewResourcesRepository() resource.Repository
}

func NewFactory() Factory {
	return &factory{}
}

type factory struct {
}

func (f *factory) NewRetrieveAllResourcesUseCase() *RetrieveAllResources {
	return &RetrieveAllResources{
		Resources: f.NewResourcesRepository(),
	}
}

func (f *factory) NewResourcesRepository() resource.Repository {
	return &inMemoryResourceRepository{}
}

type inMemoryResourceRepository struct {
}

func (resources *inMemoryResourceRepository) All() ([]resource.Resource, error) {
	return []resource.Resource{
		{
			Name: "Falco profile for Nginx",
		},
		{
			Name: "Grafana Dashboard for Traefik",
		},
	}, nil
}
