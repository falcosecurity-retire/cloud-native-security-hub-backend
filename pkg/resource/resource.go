package resource

import (
	"gopkg.in/yaml.v2"
)

type Resource struct {
	ID                ResourceID
	Version           string
	AvailableVersions []string
	Vendor            string
	Name              string
	ShortDescription  string
	Description       string
	Keywords          []string
	Icon              string
	Website           string
	Maintainers       []*Maintainer
	Rules             []*FalcoRuleData
	Policies          []*OpenPolicyAgentPolicyData
}

type Maintainer struct {
	Name string
	Link string
}

type FalcoRuleData struct {
	Raw string
}

type OpenPolicyAgentPolicyData struct {
	Raw string
}

func (r *Resource) GenerateRulesForHelmChart() []byte {
	raw := make(map[string]map[string]string)
	raw["customRules"] = map[string]string{}

	for _, rule := range r.Rules {
		raw["customRules"]["rules-"+r.ID.Slug()+".yaml"] += rule.Raw
	}

	result, _ := yaml.Marshal(raw)
	return result
}
