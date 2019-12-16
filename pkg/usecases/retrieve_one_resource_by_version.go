package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/event"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResourceByVersion struct {
	ResourceRepository resource.Repository
	EventHandler       event.Handler
	ResourceID         string
	Version            string
	Updater            resource.Updater
}

func (useCase *RetrieveOneResourceByVersion) Execute() (res *resource.Resource, err error) {
	useCase.EventHandler.Dispatch(&event.RetrievedResource{
		ResourceID: useCase.ResourceID,
		Updater:    useCase.Updater,
	})
	return useCase.ResourceRepository.FindByVersion(useCase.ResourceID, useCase.Version)
}
