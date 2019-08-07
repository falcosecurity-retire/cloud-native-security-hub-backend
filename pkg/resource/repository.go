package resource

type ResourceRepository interface {
	All() ([]Resource, error)
}
