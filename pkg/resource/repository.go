package resource

import (
	"errors"
)

type Repository interface {
	Save(*Resource) error
	FindAll() ([]*Resource, error)
	FindById(id string) (*Resource, error)
	FindByVersion(id string, version string) (*Resource, error)
}

var (
	ErrResourceNotFound = errors.New("no resource was found")
)
