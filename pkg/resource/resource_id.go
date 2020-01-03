package resource

import (
	"github.com/gosimple/slug"
)

const (
	FalcoRules              = "FalcoRules"
	OpenPolicyAgentPolicies = "OpenPolicyAgentPolicies"
)

type ResourceID struct {
	slug string
	kind string
}

func NewResourceID(name, kind string) ResourceID {
	return ResourceID{
		slug: slug.Make(name),
		kind: kind,
	}
}

func (r *ResourceID) Slug() string {
	return r.slug
}

func (r *ResourceID) Kind() string {
	return r.kind
}
