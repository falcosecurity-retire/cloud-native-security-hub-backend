package vendor

type Repository interface {
	FindAll() ([]*Vendor, error)
	FindById(id string) (*Vendor, error)
}
