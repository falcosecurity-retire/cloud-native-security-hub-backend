package resource

type Resource struct {
	ApiVersion  string
	Kind        string
	Vendor      string
	Name        string
	Description string
	Readme      string
	Keywords    []string
	Icon        string
	Maintainers []*Maintainer
	Rules       []*FalcoRuleData
	DashboardID int
}

type Maintainer struct {
	Name  string
	Email string
}

func (r *Resource) ToFalcoRule() *FalcoRule {
	return &FalcoRule{
		ApiVersion:  r.ApiVersion,
		Kind:        r.Kind,
		Vendor:      r.Vendor,
		Name:        r.Name,
		Description: r.Description,
		Readme:      r.Readme,
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
