package resource

import (
	"crypto/sha1"
	"encoding/base32"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceValidateOK(t *testing.T) {
	resource := newResource()

	assert.NoError(t, resource.Validate())
}

func TestResourceValidateKind(t *testing.T) {
	resourceWithoutKind := newResource()

	resourceWithoutKind.Kind = ""

	assert.Error(t, resourceWithoutKind.Validate())
}

func TestResourceValidateVendor(t *testing.T) {
	resourceWithoutVendor := newResource()

	resourceWithoutVendor.Vendor = ""

	assert.Error(t, resourceWithoutVendor.Validate())
}

func TestResourceValidateMaintainers(t *testing.T) {
	resourceWithoutMaintainers := newResource()

	resourceWithoutMaintainers.Maintainers = []*Maintainer{}

	assert.Error(t, resourceWithoutMaintainers.Validate())
}

func TestResourceValidateIcon(t *testing.T) {
	resourceWithoutIcon := newResource()

	resourceWithoutIcon.Icon = ""

	assert.Error(t, resourceWithoutIcon.Validate())
}

func TestResourceHash(t *testing.T) {
	resource := newResource()
	sum := sha1.Sum([]byte(string(resource.Kind) + resource.Name + resource.Vendor))
	b32 := base32.StdEncoding.EncodeToString(sum[:])
	expected := b32[:20]

	result := resource.Hash()

	assert.Equal(t, expected, result)

}

func newResource() Resource {
	return Resource{
		Kind:        "GrafanaDashboard",
		Vendor:      "Sysdig",
		Name:        "",
		Description: "",
		Rules:       nil,
		Keywords:    []string{"monitoring"},
		Icon:        "https://sysdig.com/icon.png",
		Maintainers: []*Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
		},
	}
}
