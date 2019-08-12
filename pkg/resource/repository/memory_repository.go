package repository

import "cloud-native-visibility-hub/pkg/resource"

type Memory struct {
	resources []resource.Resource
}

func (r *Memory) All() ([]resource.Resource, error) {
	return r.resources, nil
}

func (r *Memory) Add(resource resource.Resource) {
	r.resources = append(r.resources, resource)
}
