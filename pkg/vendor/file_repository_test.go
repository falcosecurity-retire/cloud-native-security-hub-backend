package vendor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepositoryWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures/vendors"
	vendorRepository, _ := FromPath(path)

	resources, _ := vendorRepository.FindAll()

	assert.Equal(t, buildResourcesFromFixtures(), resources)
}

func buildResourcesFromFixtures() []*Resource {
	resources := []*Resource{
		{
			Kind:        "Vendor",
			Name:        "Apache",
			Description: "# Apache Software Foundation\n",
			Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_Software_Foundation_Logo_%282016%29.svg/2560px-Apache_Software_Foundation_Logo_%282016%29.svg.png",
			Website:     "https://apache.org/",
		},

		{
			Kind:        "Vendor",
			Name:        "Mongo",
			Description: "# MongoDB Inc.\n",
			Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/640px-MongoDB-Logo.svg.png",
			Website:     "https://mongodb.com/",
		},
	}

	for i := range resources {
		resources[i].ID = resources[i].Hash()
	}

	return resources
}

func TestFileRepositoryReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := FromPath(nonExistentPath)

	assert.Error(t, err)
}
