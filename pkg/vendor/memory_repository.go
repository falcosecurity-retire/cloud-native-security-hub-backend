package vendor

type MemoryRepository struct {
	vendor []*Vendor
}

func NewMemoryRepository(vendor []*Vendor) Repository {
	return &MemoryRepository{
		vendor: vendor,
	}
}

func (r *MemoryRepository) FindAll() ([]*Vendor, error) {
	return r.vendor, nil
}

func (r *MemoryRepository) FindById(id string) (*Vendor, error) {
	for _, res := range r.vendor {
		if res.ID == id {
			return res, nil
		}
	}
	return nil, ErrVendorNotFound
}

func (r *MemoryRepository) Save(vendor *Vendor) error {
	r.vendor = append(r.vendor, vendor)
	return nil
}
