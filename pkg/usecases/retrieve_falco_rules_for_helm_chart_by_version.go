package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveFalcoRulesForHelmChartByVersion struct {
	ResourceRepository resource.Repository
}

func (r *RetrieveFalcoRulesForHelmChartByVersion) Execute(resourceID, version string) ([]byte, error) {
	res, err := r.ResourceRepository.FindByVersion(
		resource.NewResourceID(resourceID, resource.FalcoRules),
		version)
	if err != nil {
		return nil, err
	}
	return res.GenerateRulesForHelmChart(), nil
}
