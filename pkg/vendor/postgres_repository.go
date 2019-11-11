package vendor

import (
	"encoding/json"
	"errors"

	"database/sql"
	"database/sql/driver"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

type vendorForPostgres Vendor

func (r vendorForPostgres) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *vendorForPostgres) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &r)
}

func (r *postgresRepository) Save(vendor *Vendor) error {
	_, err := r.db.Exec(
		"INSERT INTO vendors(raw) VALUES($1)",
		vendorForPostgres(*vendor))

	return err
}

func (r *postgresRepository) FindById(id string) (*Vendor, error) {
	result := new(vendorForPostgres)
	err := r.db.QueryRow(`SELECT raw FROM vendors WHERE raw @> jsonb_build_object('id', $1::text)`, id).Scan(&result)

	if err == sql.ErrNoRows {
		return nil, ErrVendorNotFound
	}

	return (*Vendor)(result), err
}

func (r *postgresRepository) FindAll() ([]*Vendor, error) {
	rows, err := r.db.Query(`SELECT raw FROM vendors`)
	defer rows.Close()

	var result []*Vendor
	for rows.Next() {
		current := new(vendorForPostgres)
		if err = rows.Scan(&current); err != nil {
			return nil, err
		}
		result = append(result, (*Vendor)(current))
	}

	return result, err
}
