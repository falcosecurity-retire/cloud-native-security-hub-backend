package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"fmt"
	"strings"
)

type RetrieveOneResource struct {
	ResourceRepository resource.Repository
	Hash               string
}

func (useCase *RetrieveOneResource) Execute() (res *resource.Resource, err error) {
	resources, err := useCase.ResourceRepository.FindAll()
	if err != nil {
		return
	}

	for _, res := range resources {
		resourceHash := strings.ToLower(res.Hash())
		hashToLookFor := strings.ToLower(useCase.Hash)
		if resourceHash == hashToLookFor {
			return res, nil
		}
	}

	err = fmt.Errorf("not found")

	return
}