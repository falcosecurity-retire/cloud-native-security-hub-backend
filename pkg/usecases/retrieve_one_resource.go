package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/event"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveOneResource struct {
	ResourceRepository resource.Repository
	ResourceID         string
	EventDispatcher    event.Dispatcher
	Updater            resource.Updater
}

func (useCase *RetrieveOneResource) Execute() (res *resource.Resource, err error) {
	useCase.EventDispatcher.Dispatch(&event.RetrievedResource{
		ResourceID: useCase.ResourceID,
		Updater:    useCase.Updater,
	})
	return useCase.ResourceRepository.FindById(useCase.ResourceID)
}
