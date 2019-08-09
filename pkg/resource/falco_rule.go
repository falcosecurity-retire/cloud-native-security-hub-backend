package resource

type FalcoRule struct {
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
}

type FalcoRuleData struct {
	Raw  string `json:"raw" yaml:"raw"`
}
