package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveFalcoRulesForHelmChart struct {
	ResourceRepository resource.Repository
	ResourceID         string
}

func (useCase *RetrieveFalcoRulesForHelmChart) Execute() ([]byte, error) {
	res, err := useCase.ResourceRepository.FindById(useCase.ResourceID)
	if err != nil {
		return nil, err
	}
	return res.GenerateRulesForHelmChart(), nil
}
