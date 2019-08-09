package resource

import (
	"fmt"
	"strings"
)

type Resource struct {
	ApiVersion  string           `json:"apiVersion" yaml:"apiVersion"`
	Kind        string           `json:"kind" yaml:"kind"`
	Vendor      string           `json:"vendor" yaml:"vendor"`
	Name        string           `json:"name" yaml:"name"`
	Description string           `json:"description" yaml:"description"`
	Keywords    []string         `json:"keywords" yaml:"keywords"`
	Icon        string           `json:"icon" yaml:"icon"`
	Maintainers []*Maintainer    `json:"maintainers" yaml:"maintainers"`
	Rules       []*FalcoRuleData `json:"rules" yaml:"rules"`
	DashboardID int
}

type Maintainer struct {
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
}

func (r *Resource) ToFalcoRule() *FalcoRule {
	return &FalcoRule{
		ApiVersion:  r.ApiVersion,
		Kind:        r.Kind,
		Vendor:      r.Vendor,
		Name:        r.Name,
		Description: r.Description,
		Keywords:    r.Keywords,
		Icon:        r.Icon,
		Maintainers: r.Maintainers,
		Rules:       r.Rules,
	}
}

func (r *Resource) ToGrafanaDashboard() *GrafanaDashboard {
	return &GrafanaDashboard{
		ApiVersion:  r.ApiVersion,
		Kind:        r.Kind,
		Vendor:      r.Vendor,
		Keywords:    r.Keywords,
		Icon:        r.Icon,
		Maintainers: r.Maintainers,
		DashboardID: r.DashboardID,
	}
}

func (r *Resource) Validate() error {
	var errors []string

	if r.Kind == "" {
		errors = append(errors, "the resource must have a defined Kind")
	}
	if r.ApiVersion == "" {
		errors = append(errors, "the resource does not have an API Version")
	}
	if r.Vendor == "" {
		errors = append(errors, "the resource must be assigned to a vendor")
	}
	if len(r.Maintainers) == 0 {
		errors = append(errors, "the resource must have at least one maintainer")
	}
	if r.Icon == "" {
		errors = append(errors, "the resource must have a valid icon")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ","))
	}

	return nil
}
