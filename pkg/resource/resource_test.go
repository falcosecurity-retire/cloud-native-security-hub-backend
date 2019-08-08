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

	assert.Equal(t, resource.ToFalcoRule(), rule)
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

	assert.Equal(t, resource.ToGrafanaDashboard(), dashboard)
}
