package usecases

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrieveOneRawResource struct {
	ResourceRepository resource.Repository
	Hash               string
}

func (useCase *RetrieveOneRawResource) Execute() (raw resource.Raw, err error) {
	res, err := useCase.ResourceRepository.FindById(useCase.Hash)
	if err != nil {
		return
	}
	return res, nil
}
