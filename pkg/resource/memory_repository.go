package resource

import (
	"fmt"
	"strings"
)

type MemoryRepository struct {
	resources []*Resource
}

func NewMemoryRepository(resources []*Resource) Repository {
	return &MemoryRepository{
		resources: resources,
	}
}

func (r *MemoryRepository) FindAll() ([]*Resource, error) {
	return r.resources, nil
}

func (r *MemoryRepository) FindById(id string) (*Resource, error) {
	idToFind := strings.ToLower(id)
	for _, res := range r.resources {
		if res.ID == idToFind {
			return res, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r *MemoryRepository) Add(resource Resource) {
	r.resources = append(r.resources, &resource)
}
