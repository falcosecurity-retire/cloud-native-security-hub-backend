package resource

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

type resourceForPostgres Resource

func (r resourceForPostgres) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *resourceForPostgres) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, &r)
}

func (r *postgresRepository) Save(resource *Resource) error {
	_, err := r.db.Exec(
		"INSERT INTO security_resources(raw) VALUES($1)",
		resourceForPostgres(*resource))

	return err
}

func (r *postgresRepository) FindById(id string) (*Resource, error) {
	result := new(resourceForPostgres)
	err := r.db.QueryRow(`SELECT raw FROM security_resources WHERE raw @> jsonb_build_object('id', $1::text)`, id).Scan(&result)

	return (*Resource)(result), err
}

func (r *postgresRepository) FindAll() ([]*Resource, error) {
	rows, err := r.db.Query(`SELECT raw FROM security_resources`)
	defer rows.Close()

	var result []*Resource
	for rows.Next() {
		current := new(resourceForPostgres)
		if err = rows.Scan(&current); err != nil {
			return nil, err
		}
		result = append(result, (*Resource)(current))
	}

	return result, err
}
