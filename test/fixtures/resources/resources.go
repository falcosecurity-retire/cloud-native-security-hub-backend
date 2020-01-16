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

func TrustedRegistriesImagesWithoutAvailableVersions() *resource.Resource {
	return &resource.Resource{
		Vendor:           "Kubernetes",
		ID:               resource.NewResourceID("Kubernetes Trusted Registry Images", resource.OpenPolicyAgentPolicies),
		Name:             "Kubernetes Trusted Registry Images",
		ShortDescription: "Ensure all images running in the Kubernetes cluster comes from a trusted registry",
		Version:          "0.1.0",
		Description:      "# Kubernetes Trusted Registry Images\n",
		Keywords:         []string{"kubernetes"},
		Icon:             "https://upload.wikimedia.org/wikipedia/commons/thumb/3/39/Kubernetes_logo_without_workmark.svg/300px-Kubernetes_logo_without_workmark.svg.png",
		Maintainers: []*resource.Maintainer{
			{
				Name: "Néstor Salceda",
				Link: "https://github.com/nestorsalceda",
			},
		},
		Policies: []*resource.OpenPolicyAgentPolicyData{
			{
				Raw: "package kubernetes.admission\n\ndeny[msg] {}\n",
			},
		},
	}
}

func TrustedRegistriesImages() *resource.Resource {
	result := TrustedRegistriesImagesWithoutAvailableVersions()
	result.AvailableVersions = []string{"0.1.0"}

	return result
}
