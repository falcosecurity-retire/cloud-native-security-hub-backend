package resource

type GrafanaDashboard struct {
	ApiVersion  string
	Kind        string
	Vendor      string
	Keywords    []string
	Icon        string
	Maintainers []*Maintainer
	DashboardID int
}
