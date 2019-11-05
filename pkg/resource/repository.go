package resource

type Repository interface {
	Save(*Resource) error
	FindAll() ([]*Resource, error)
	FindById(id string) (*Resource, error)
}
