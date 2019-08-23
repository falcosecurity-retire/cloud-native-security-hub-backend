package vendor

type Repository interface {
	FindAll() ([]*Resource, error)
	FindById(id string) (*Resource, error)
}
