package resources

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

func Apache() *resource.Resource {
	result := ApacheWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

func ApacheWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		ID:          resource.NewResourceID("apache", resource.FalcoRules),
		Vendor:      "Apache",
		Name:        "Apache",
		Version:     "1.0.0",
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
				Raw: "- macro: apache_consider_syscalls\n  condition: (evt.num < 0)\n",
			},
		},
	}
}

func MongoDB() *resource.Resource {
	result := MongoDBWithoutAvailableVersions()
	result.AvailableVersions = []string{"1.0.0"}

	return result
}

func MongoDBWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		Vendor:      "Mongo",
		ID:          resource.NewResourceID("mongodb", resource.FalcoRules),
		Name:        "MongoDB",
		Version:     "1.0.0",
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
				Raw: "- macro: mongo_consider_syscalls\n  condition: (evt.num < 0)\n",
			},
		},
	}
}
