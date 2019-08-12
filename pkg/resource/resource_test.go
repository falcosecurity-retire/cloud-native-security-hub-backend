package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceValidateOK(t *testing.T) {
	resource := newResource()

	assert.NoError(t, resource.Validate())
}

func TestResourceValidateApiVersion(t *testing.T) {
	resourceWithoutVersion := newResource()

	resourceWithoutVersion.ApiVersion = ""

	assert.Error(t, resourceWithoutVersion.Validate())
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

func newResource() Resource {
	return Resource{
		ApiVersion:  "v1",
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
		DashboardID: 0,
	}
}
