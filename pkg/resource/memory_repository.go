package resource

import (
	"fmt"
	"strings"
)

type Memory struct {
	resources []*Resource
}

func (r *Memory) FindAll() ([]*Resource, error) {
	return r.resources, nil
}

func (r *Memory) FindById(id string) (*Resource, error) {
	idToFind := strings.ToLower(id)
	for _, res := range r.resources {
		if res.ID == idToFind {
			return res, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r *Memory) Add(resource Resource) {
	r.resources = append(r.resources, &resource)
}
