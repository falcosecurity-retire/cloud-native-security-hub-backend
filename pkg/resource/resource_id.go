package resource

import (
	"github.com/gosimple/slug"
)

type ResourceID struct {
	name string
	kind string
}

func NewResourceID(name, kind string) ResourceID {
	return ResourceID{
		name: slug.Make(name),
		kind: kind,
	}
}
