package resource

type Repository interface {
	All() ([]Resource, error)
}
