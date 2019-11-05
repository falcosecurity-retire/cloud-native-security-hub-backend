package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

func TestGetResourcesFromPathWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures/resources"
	resources, _ := GetResourcesFromPath(path)

	assert.Equal(t, []*resource.Resource{
		{
			ID:          "apache",
			Kind:        "FalcoRules",
			Vendor:      "Apache",
			Name:        "Apache",
			Description: "# Apache Falco Rules\n",
			Keywords:    []string{"web"},
			Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_HTTP_server_logo_%282016%29.svg/300px-Apache_HTTP_server_logo_%282016%29.svg.png",
			Maintainers: []*resource.Maintainer{
				{
					Name: "nestorsalceda",
					Link: "github.com/nestorsalceda",
				},
				{
					Name: "fedebarcelona",
					Link: "github.com/tembleking",
				},
			},
			Rules: []*resource.FalcoRuleData{
				{
					Raw: "- macro: apache_consider_syscalls\n  condition: (evt.num < 0)",
				},
			},
		},

		{
			Kind:        "FalcoRules",
			Vendor:      "Mongo",
			ID:          "mongodb",
			Name:        "MongoDB",
			Description: "# MongoDB Falco Rules\n",
			Keywords:    []string{"database"},
			Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/2560px-MongoDB-Logo.svg.png",
			Maintainers: []*resource.Maintainer{
				{
					Name: "nestorsalceda",
					Link: "github.com/nestorsalceda",
				},
				{
					Name: "fedebarcelona",
					Link: "github.com/tembleking",
				},
			},
			Rules: []*resource.FalcoRuleData{
				{
					Raw: "- macro: mongo_consider_syscalls\n  condition: (evt.num < 0)",
				},
			},
		},
	}, resources)
}

func TestGetResourcesFromPathReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := GetResourcesFromPath(nonExistentPath)

	assert.Error(t, err)
}
