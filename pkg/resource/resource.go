package resource

import (
	"crypto/sha1"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

type Kind string

const (
	FALCO_RULE Kind = "FalcoRule"
	VENDOR     Kind = "Vendor"
)

type Resource struct {
	ID          string           `json:"id,omitempty" yaml:"id,omitempty"`
	ApiVersion  string           `json:"apiVersion" yaml:"apiVersion"`
	Kind        Kind             `json:"kind" yaml:"kind"`
	Vendor      string           `json:"vendor" yaml:"vendor"`
	Name        string           `json:"name" yaml:"name"`
	Description string           `json:"description" yaml:"description"`
	Keywords    []string         `json:"keywords" yaml:"keywords"`
	Icon        string           `json:"icon" yaml:"icon"`
	Website     string           `json:"website" yaml:"website"`
	Maintainers []*Maintainer    `json:"maintainers" yaml:"maintainers"`
	Rules       []*FalcoRuleData `json:"rules" yaml:"rules"`
	DashboardID int
}

type resourceAlias Resource // Avoid stack overflow while marshalling / unmarshalling

func (r *Resource) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	res := resourceAlias{}
	err = unmarshal(&res)
	if err != nil {
		return
	}
	*r = Resource(res)
	r.ID = r.Hash()
	return
}

func (r *Resource) MarshalYAML() (interface{}, error) {
	x := resourceAlias(*r)
	x.ID = r.Hash()
	return yaml.Marshal(x)
}

func (r *Resource) UnmarshalJSON(data []byte) (err error) {
	res := resourceAlias{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	*r = Resource(res)
	r.ID = r.Hash()
	return
}

func (r *Resource) MarshalJSON() ([]byte, error) {
	x := resourceAlias(*r)
	x.ID = r.Hash()
	return json.Marshal(x)
}

type Maintainer struct {
	Name  string `json:"name" yaml:"name"`
	Email string `json:"email" yaml:"email"`
}

func (r *Resource) Validate() error {
	var errors []string

	if r.Kind == "" {
		errors = append(errors, "the resource must have a defined Kind")
	}
	if r.ApiVersion == "" {
		errors = append(errors, "the resource does not have an API Version")
	}
	if r.Kind != VENDOR && r.Vendor == "" {
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

func (r *Resource) Hash() string {
	sum := sha1.Sum([]byte(r.ApiVersion + string(r.Kind) + r.Name + r.Vendor))
	b32 := base32.StdEncoding.EncodeToString(sum[:])
	return b32[:20]
}
