package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/event"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type RetrieveFalcoRulesForHelmChart struct {
	ResourceRepository resource.Repository
	EventHandler       event.Handler
	Updater            resource.Updater
	ResourceID         string
}

func (useCase *RetrieveFalcoRulesForHelmChart) Execute() ([]byte, error) {
	res, err := useCase.ResourceRepository.FindById(useCase.ResourceID)
	if err != nil {
		return nil, err
	}
	useCase.EventHandler.Dispatch(&event.RetrievedResource{
		ResourceID: useCase.ResourceID,
		Updater:    useCase.Updater,
	})
	return res.GenerateRulesForHelmChart(), nil
}
