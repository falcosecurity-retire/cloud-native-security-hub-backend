package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
)

func TestGetResourcesFromPathWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures/resources"
	parsed, _ := GetResourcesFromPath(path)

	assert.Equal(t, []*resource.Resource{
		resources.Apache(),
		resources.MongoDB(),
	}, parsed)
}

func TestGetResourcesFromPathReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := GetResourcesFromPath(nonExistentPath)

	assert.Error(t, err)
}
