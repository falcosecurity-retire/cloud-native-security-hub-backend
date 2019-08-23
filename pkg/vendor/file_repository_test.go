package vendor

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepositoryWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures/vendors"
	vendorRepository, _ := NewFile(path)

	resources, _ := vendorRepository.FindAll()

	assert.Equal(t, buildResourcesFromFixtures(), resources)
}

func buildResourcesFromFixtures() []*resource.Resource {
	resources := []*resource.Resource{
		{
			Kind:        "Vendor",
			Name:        "Apache",
			Description: "# Apache Software Foundation\n",
			Keywords:    []string{"web"},
			Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_Software_Foundation_Logo_%282016%29.svg/2560px-Apache_Software_Foundation_Logo_%282016%29.svg.png",
			Website:     "https://apache.org/",
			Maintainers: []*resource.Maintainer{
				{
					Name:  "nestorsalceda",
					Email: "nestor.salceda@sysdig.com",
				},
				{
					Name:  "fedebarcelona",
					Email: "fede.barcelona@sysdig.com",
				},
			},
		},

		{
			Kind:        "Vendor",
			Name:        "Mongo",
			Description: "# MongoDB Inc.\n",
			Keywords:    []string{"database"},
			Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/640px-MongoDB-Logo.svg.png",
			Website:     "https://mongodb.com/",
			Maintainers: []*resource.Maintainer{
				{
					Name:  "nestorsalceda",
					Email: "nestor.salceda@sysdig.com",
				},
				{
					Name:  "fedebarcelona",
					Email: "fede.barcelona@sysdig.com",
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
