package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepositoryWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures/resources"
	fileRepository, _ := NewFile(path)

	resources, _ := fileRepository.FindAll()

	assert.Equal(t, buildResourcesFromFixtures(), resources)
}

func buildResourcesFromFixtures() []*Resource {
	resources := []*Resource{
		{
			Kind:        "FalcoRules",
			Vendor:      "Apache",
			Name:        "Apache",
			Description: "# Apache Falco Rules\n",
			Keywords:    []string{"web"},
			Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_HTTP_server_logo_%282016%29.svg/300px-Apache_HTTP_server_logo_%282016%29.svg.png",
			Maintainers: []*Maintainer{
				{
					Name:  "nestorsalceda",
					Email: "nestor.salceda@sysdig.com",
				},
				{
					Name:  "fedebarcelona",
					Email: "fede.barcelona@sysdig.com",
				},
			},
			Rules: []*FalcoRuleData{
				{
					Raw: "- macro: apache_consider_syscalls\n  condition: (evt.num < 0)",
				},
			},
		},

		{
			Kind:        "FalcoRules",
			Vendor:      "Mongo",
			Name:        "MongoDB",
			Description: "# MongoDB Falco Rules\n",
			Keywords:    []string{"database"},
			Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/2560px-MongoDB-Logo.svg.png",
			Maintainers: []*Maintainer{
				{
					Name:  "nestorsalceda",
					Email: "nestor.salceda@sysdig.com",
				},
				{
					Name:  "fedebarcelona",
					Email: "fede.barcelona@sysdig.com",
				},
			},
			Rules: []*FalcoRuleData{
				{
					Raw: "- macro: mongo_consider_syscalls\n  condition: (evt.num < 0)",
				},
			},
		},
	}

	for i := range resources {
		resources[i].ID = resources[i].Hash()
	}

	return resources
}

func TestFileRepositoryReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := NewFile(nonExistentPath)

	assert.Error(t, err)
}
