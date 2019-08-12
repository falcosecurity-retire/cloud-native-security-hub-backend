package resource

type Memory struct {
	resources []Resource
}

func (r *Memory) All() ([]Resource, error) {
	return r.resources, nil
}

func (r *Memory) Add(resource Resource) {
	r.resources = append(r.resources, resource)
}
