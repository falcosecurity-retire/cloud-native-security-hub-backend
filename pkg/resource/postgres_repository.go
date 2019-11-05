package resource

import (
	"encoding/json"
	"errors"

	"database/sql"
	"database/sql/driver"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
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

func (r *PostgresRepository) Save(resource *Resource) error {
	_, err := r.db.Exec(
		"INSERT INTO security_resources(raw) VALUES($1)",
		resourceForPostgres(*resource))

	return err
}

func (r *PostgresRepository) FindById(id string) (*Resource, error) {
	result := new(resourceForPostgres)
	err := r.db.QueryRow(`SELECT raw FROM security_resources WHERE raw @> jsonb_build_object('id', $1::text)`, id).Scan(&result)

	return (*Resource)(result), err
}
