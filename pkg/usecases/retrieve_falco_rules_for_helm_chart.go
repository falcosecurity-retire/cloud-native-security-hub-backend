package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveFalcoRulesForHelmChart struct {
	ResourceRepository resource.Repository
}

func (r *RetrieveFalcoRulesForHelmChart) Execute(resourceID string) ([]byte, error) {
	res, err := r.ResourceRepository.FindById(
		resource.NewResourceID(resourceID, resource.FalcoRules))

	if err != nil {
		return nil, err
	}
	return res.GenerateRulesForHelmChart(), nil
}
