package vendor

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

type Kind string

const (
	VENDOR Kind = "Vendor"
)

type Vendor struct {
	ID          string `json:"id,omitempty" yaml:"-"`
	Kind        Kind   `json:"kind" yaml:"kind"`
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Icon        string `json:"icon" yaml:"icon"`
	Website     string `json:"website" yaml:"website"`
}

type vendorAlias Vendor // Avoid stack overflow while marshalling / unmarshalling

func (r *Vendor) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	res := vendorAlias{}
	err = unmarshal(&res)
	if err != nil {
		return
	}
	*r = Vendor(res)
	r.ID = r.generateID()
	return
}

func (r *Vendor) MarshalYAML() (interface{}, error) {
	x := vendorAlias(*r)
	r.ID = r.generateID()
	return yaml.Marshal(x)
}

func (r *Vendor) UnmarshalJSON(data []byte) (err error) {
	res := vendorAlias{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	*r = Vendor(res)
	r.ID = r.generateID()
	return
}

func (r *Vendor) MarshalJSON() ([]byte, error) {
	x := vendorAlias(*r)
	r.ID = r.generateID()
	return json.Marshal(x)
}

func (r *Vendor) Validate() error {
	var errors []string

	if r.Kind == "" {
		errors = append(errors, "the vendor must have a defined Kind")
	}

	if r.Icon == "" {
		errors = append(errors, "the vendor must have a valid icon")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, ","))
	}

	return nil
}

func (r *Vendor) generateID() string {
	return strings.ToLower(r.Name)
}
