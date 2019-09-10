package vendor

import (
	"fmt"
	"strings"
)

type MemoryRepository struct {
	vendor []*Vendor
}

func NewMemoryRepository(vendor []*Vendor) Repository {
	return &MemoryRepository{
		vendor: vendor,
	}
}

func (r *MemoryRepository) FindAll() ([]*Vendor, error) {
	return r.vendor, nil
}

func (r *MemoryRepository) FindById(id string) (*Vendor, error) {
	idToFind := strings.ToLower(id)
	for _, res := range r.vendor {
		if res.ID == idToFind {
			return res, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (r *MemoryRepository) Add(vendor Vendor) {
	r.vendor = append(r.vendor, &vendor)
}
