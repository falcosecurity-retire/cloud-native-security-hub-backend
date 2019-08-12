package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestResource_Validate(t *testing.T) {
	fullResource := Resource{
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

	resourceWithoutVersion := fullResource
	resourceWithoutVersion.ApiVersion = ""
	assert.Error(t, resourceWithoutVersion.Validate())

	resourceWithoutKind := fullResource
	resourceWithoutKind.Kind = ""
	assert.Error(t, resourceWithoutKind.Validate())

	resourceWithoutVendor := fullResource
	resourceWithoutVendor.Kind = ""
	assert.Error(t, resourceWithoutVendor.Validate())

	resourceWithoutMaintainers := fullResource
	resourceWithoutMaintainers.Maintainers = []*Maintainer{}
	assert.Error(t, resourceWithoutMaintainers.Validate())

	resourceWithoutIcon := fullResource
	resourceWithoutIcon.Kind = ""
	assert.Error(t, resourceWithoutIcon.Validate())

	assert.Equal(t, nil, fullResource.Validate())
}
