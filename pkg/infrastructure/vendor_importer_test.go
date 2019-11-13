package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
	"github.com/falcosecurity/cloud-native-security-hub/test/fixtures/vendors"
)

func TestGetVendorsWalksADirectoryAndExtractVendors(t *testing.T) {
	path := "../../test/fixtures/vendors"
	parsed, _ := GetVendorsFromPath(path)

	assert.Equal(t, []*vendor.Vendor{
		vendors.Apache(),
		vendors.Mongo(),
	}, parsed)
}

func TestGetVendorsReturnsAnErrorIfPathDoesNotExist(t *testing.T) {
	nonExistentPath := "../foo"

	_, err := GetVendorsFromPath(nonExistentPath)

	assert.Error(t, err)
}
