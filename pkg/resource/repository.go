package resource

type Repository interface {
	FindAll() ([]*Resource, error)
	FindById(id string) (*Resource, error)
}
