package vendor

type Repository interface {
	Save(*Vendor) error

	FindAll() ([]*Vendor, error)
	FindById(id string) (*Vendor, error)
}
