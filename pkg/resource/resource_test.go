package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResource_ToFalcoRule(t *testing.T) {
	resource := Resource{
		ApiVersion:  "v1",
		Kind:        "FalcoRule",
		Vendor:      "Sysdig",
		Name:        "Foo",
		Description: "FooBar",
		Keywords:    []string{"monitoring", "security"},
		Icon:        "https://sysdig.com/icon.png",
		Maintainers: []*Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
			{
				Name:  "nestor",
				Email: "nestor.salceda@sysdig.com",
			},
		},
		Rules: []*FalcoRuleData{
			{
				Raw: "testRule",
			},
		},
		DashboardID: 0,
	}

	rule := &FalcoRule{
		ApiVersion:  "v1",
		Kind:        "FalcoRule",
		Vendor:      "Sysdig",
		Name:        "Foo",
		Description: "FooBar",
		Keywords:    []string{"monitoring", "security"},
		Icon:        "https://sysdig.com/icon.png",
		Maintainers: []*Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
			{
				Name:  "nestor",
				Email: "nestor.salceda@sysdig.com",
			},
		},
		Rules: []*FalcoRuleData{
			{
				Raw: "testRule",
			},
		},
	}

	assert.Equal(t, rule, resource.ToFalcoRule())
}

func TestResource_ToGrafanaDashboard(t *testing.T) {
	dashboard := &GrafanaDashboard{
		ApiVersion: "v1",
		Kind:       "GrafanaDashboard",
		Vendor:     "Sysdig",
		Keywords:   []string{"monitoring", "security"},
		Icon:       "https://sysdig.com/icon.png",
		Maintainers: []*Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
			{
				Name:  "nestor",
				Email: "nestor.salceda@sysdig.com",
			},
		},
		DashboardID: 0,
	}

	resource := Resource{
		ApiVersion:  "v1",
		Kind:        "GrafanaDashboard",
		Vendor:      "Sysdig",
		Name:        "",
		Description: "",
		Rules:       nil,
		Keywords:    []string{"monitoring", "security"},
		Icon:        "https://sysdig.com/icon.png",
		Maintainers: []*Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
			{
				Name:  "nestor",
				Email: "nestor.salceda@sysdig.com",
			},
		},
		DashboardID: 0,
	}

	assert.Equal(t, dashboard, resource.ToGrafanaDashboard())
}

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
