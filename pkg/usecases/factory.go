package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
	"cloud-native-visibility-hub/pkg/resource/repository"
	"log"
	"os"
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
		ResourceRepository: f.NewResourcesRepository(),
	}
}

func (f *factory) NewResourcesRepository() resource.Repository {
	resourcesPath, ok := os.LookupEnv("RESOURCES_PATH")
	if !ok {
		log.Println("The RESOURCES_PATH env var is not set")
		os.Exit(1)
	}
	return &repository.File{Path: resourcesPath}
}
