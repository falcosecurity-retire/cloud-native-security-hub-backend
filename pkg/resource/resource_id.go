package resource

import (
	"github.com/gosimple/slug"
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
