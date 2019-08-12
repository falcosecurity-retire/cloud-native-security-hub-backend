package usecases

import (
	"cloud-native-visibility-hub/pkg/resource"
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
	repo, err := resource.NewFile(resourcesPath)
	if err != nil {
		log.Println("the resource repository of type file does not exist")
		os.Exit(1)
	}
	return repo
}
