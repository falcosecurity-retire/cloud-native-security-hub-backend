package infrastructure_test

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/infrastructure"
	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"

	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetVendorsWalksADirectoryAndExtractVendors(t *testing.T) {
	path := "../../test/fixtures/vendors"
	parsed, _ := infrastructure.GetVendorsFromPath(path)

	assert.Equal(t, []*vendor.Vendor{
		vendors.Apache(),
		vendors.Mongo(),
	}, parsed)
}

func TestGetVendorsReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := infrastructure.GetVendorsFromPath(nonExistentPath)

	assert.Error(t, err)
}
