package infrastructure_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/infrastructure"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResourcesFromPathWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures/resources"
	parsed, _ := infrastructure.GetResourcesFromPath(path)

	assert.Equal(t, []*resource.Resource{
		resources.Apache(),
		resources.MongoDB(),
	}, parsed)
}

func TestGetResourcesFromPathReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := infrastructure.GetResourcesFromPath(nonExistentPath)

	assert.Error(t, err)
}
