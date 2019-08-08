package resource

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
