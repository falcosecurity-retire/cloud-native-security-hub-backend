package resource

import (
	"encoding/json"

	"github.com/gosimple/slug"
	"gopkg.in/yaml.v2"
)

type Resource struct {
	ID                string           `json:"id,omitempty" yaml:"id,omitempty"`
	Kind              string           `json:"kind" yaml:"kind"`
	Version           string           `json:"version" yaml:"version"`
	AvailableVersions []string         `json:"availableVersions" yaml:"-"`
	Vendor            string           `json:"vendor" yaml:"vendor"`
	Name              string           `json:"name" yaml:"name"`
	ShortDescription  string           `json:"shortDescription" yaml:"shortDescription"`
	Description       string           `json:"description" yaml:"description"`
	Keywords          []string         `json:"keywords" yaml:"keywords"`
	Icon              string           `json:"icon" yaml:"icon"`
	Website           string           `json:"website" yaml:"website"`
	Maintainers       []*Maintainer    `json:"maintainers" yaml:"maintainers"`
	Rules             []*FalcoRuleData `json:"rules" yaml:"rules"`
}

func (r *Resource) GenerateRulesForHelmChart() []byte {
	raw := make(map[string]map[string]string)
	raw["customRules"] = map[string]string{}

	for _, rule := range r.Rules {
		raw["customRules"]["rules-"+r.ID+".yaml"] += rule.Raw
	}

	result, _ := yaml.Marshal(raw)
	return result
}

type resourceAlias Resource // Avoid stack overflow while marshalling / unmarshalling

func (r *Resource) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	res := resourceAlias{}
	err = unmarshal(&res)
	if err != nil {
		return
	}
	*r = Resource(res)
	r.ID = r.generateID()
	return
}

func (r *Resource) MarshalYAML() (interface{}, error) {
	x := resourceAlias(*r)
	x.ID = r.generateID()
	return yaml.Marshal(x)
}

func (r *Resource) UnmarshalJSON(data []byte) (err error) {
	res := resourceAlias{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	*r = Resource(res)
	r.ID = r.generateID()
	return
}

func (r *Resource) MarshalJSON() ([]byte, error) {
	x := resourceAlias(*r)
	x.ID = r.generateID()
	return json.Marshal(x)
}

type Maintainer struct {
	Name string `json:"name" yaml:"name"`
	Link string `json:"link" yaml:"link"`
}

type FalcoRuleData struct {
	Raw string `json:"raw" yaml:"raw"`
}

func (r *Resource) generateID() string {
	return slug.Make(r.Name)
}
