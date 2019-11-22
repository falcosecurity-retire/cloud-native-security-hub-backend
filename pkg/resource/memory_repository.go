package resource

import (
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
	return nil, ErrResourceNotFound
}

func (r *MemoryRepository) Save(resource *Resource) error {
	r.resources = append(r.resources, resource)
	return nil
}

func (r *MemoryRepository) FindByVersion(id string, version string) (*Resource, error) {
	for _, res := range r.resources {
		if res.ID == strings.ToLower(id) && res.Version == version {
			return res, nil
		}
	}
	return nil, ErrResourceNotFound
}
