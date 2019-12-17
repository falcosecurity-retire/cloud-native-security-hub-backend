package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveFalcoRulesForHelmChartByVersion struct {
	ResourceRepository resource.Repository
	ResourceID         string
	Version            string
}

func (useCase *RetrieveFalcoRulesForHelmChartByVersion) Execute() ([]byte, error) {
	res, err := useCase.ResourceRepository.FindByVersion(useCase.ResourceID, useCase.Version)
	if err != nil {
		return nil, err
	}
	return res.GenerateRulesForHelmChart(), nil
}
