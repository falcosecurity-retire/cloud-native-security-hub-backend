package resource

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type ResourceDTO struct {
	ID                string                          `json:"id" yaml:"-"`
	Kind              string                          `json:"kind" yaml:"kind"`
	Version           string                          `json:"version" yaml:"version"`
	AvailableVersions []string                        `json:"availableVersions" yaml:"-"`
	Vendor            string                          `json:"vendor" yaml:"vendor"`
	Name              string                          `json:"name" yaml:"name"`
	ShortDescription  string                          `json:"shortDescription" yaml:"shortDescription"`
	Description       string                          `json:"description" yaml:"description"`
	Keywords          []string                        `json:"keywords" yaml:"keywords"`
	Icon              string                          `json:"icon" yaml:"icon"`
	Website           string                          `json:"website" yaml:"website"`
	Maintainers       []*MaintainerDTO                `json:"maintainers" yaml:"maintainers"`
	Rules             []*FalcoRuleDataDTO             `json:"rules,omitempty" yaml:"rules"`
	Policies          []*OpenPolicyAgentPolicyDataDTO `json:"policies,omitempty" yaml:"policies"`
}

type MaintainerDTO struct {
	Name string `json:"name" yaml:"name"`
	Link string `json:"link" yaml:"link"`
}

type FalcoRuleDataDTO struct {
	Raw string `json:"raw" yaml:"raw"`
}

type OpenPolicyAgentPolicyDataDTO struct {
	Raw string `json:"raw" yaml:"raw"`
}

func NewResourceDTO(entity *Resource) *ResourceDTO {
	return &ResourceDTO{
		ID:                entity.ID.Slug(),
		Kind:              entity.ID.Kind(),
		Version:           entity.Version,
		AvailableVersions: entity.AvailableVersions,
		Vendor:            entity.Vendor,
		Name:              entity.Name,
		ShortDescription:  entity.ShortDescription,
		Description:       entity.Description,
		Keywords:          entity.Keywords,
		Icon:              entity.Icon,
		Website:           entity.Website,
		Maintainers:       parseMaintainers(entity.Maintainers),
		Rules:             parseRules(entity.Rules),
		Policies:          parsePolicies(entity.Policies),
	}
}

func parseMaintainers(maintainers []*Maintainer) []*MaintainerDTO {
	var result []*MaintainerDTO

	for _, maintainer := range maintainers {
		result = append(result, &MaintainerDTO{
			Name: maintainer.Name,
			Link: maintainer.Link,
		})
	}

	return result
}

func parseRules(rules []*FalcoRuleData) []*FalcoRuleDataDTO {
	var result []*FalcoRuleDataDTO

	for _, rule := range rules {
		result = append(result, &FalcoRuleDataDTO{
			Raw: rule.Raw,
		})
	}

	return result
}

func parsePolicies(policies []*OpenPolicyAgentPolicyData) []*OpenPolicyAgentPolicyDataDTO {
	var result []*OpenPolicyAgentPolicyDataDTO

	for _, policy := range policies {
		result = append(result, &OpenPolicyAgentPolicyDataDTO{
			Raw: policy.Raw,
		})
	}

	return result
}

func (r *ResourceDTO) ToEntity() *Resource {
	return &Resource{
		ID:                NewResourceID(r.Name, r.Kind),
		Version:           r.Version,
		AvailableVersions: r.AvailableVersions,
		Vendor:            r.Vendor,
		Name:              r.Name,
		ShortDescription:  r.ShortDescription,
		Description:       r.Description,
		Keywords:          r.Keywords,
		Icon:              r.Icon,
		Website:           r.Website,
		Maintainers:       toEntityMaintainers(r.Maintainers),
		Rules:             toEntityFalcoRuleData(r.Rules),
		Policies:          toEntityOpenPolicyAgentPolicyData(r.Policies),
	}
}

func toEntityMaintainers(maintainers []*MaintainerDTO) []*Maintainer {
	var result []*Maintainer

	for _, maintainer := range maintainers {
		result = append(result, &Maintainer{
			Name: maintainer.Name,
			Link: maintainer.Link,
		})
	}

	return result
}

func toEntityFalcoRuleData(rules []*FalcoRuleDataDTO) []*FalcoRuleData {
	var result []*FalcoRuleData

	for _, rule := range rules {
		result = append(result, &FalcoRuleData{
			Raw: rule.Raw,
		})
	}

	return result
}

func toEntityOpenPolicyAgentPolicyData(policies []*OpenPolicyAgentPolicyDataDTO) []*OpenPolicyAgentPolicyData {
	var result []*OpenPolicyAgentPolicyData

	for _, policy := range policies {
		result = append(result, &OpenPolicyAgentPolicyData{
			Raw: policy.Raw,
		})
	}

	return result
}

func (r ResourceDTO) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ResourceDTO) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &r)
}
