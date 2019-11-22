package vendor

import (
	"errors"
)

var (
	ErrVendorNotFound = errors.New("no vendor was found")
)

type Repository interface {
	Save(*Vendor) error

	FindAll() ([]*Vendor, error)
	FindById(id string) (*Vendor, error)
}
