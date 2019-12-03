package resource

import (
	"encoding/json"
	"errors"

	"database/sql"
	"database/sql/driver"
	"github.com/lib/pq"

	"github.com/Masterminds/semver"
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
	transaction, err := r.db.Begin()
	_, err = transaction.Exec("INSERT INTO security_resources(raw) VALUES($1)", resourceForPostgres(*resource))

	if existent := retrieveExistingVersion(transaction, resource); existent == "" {
		_, err = transaction.Exec(
			"INSERT INTO latest_security_resources(raw) VALUES($1)",
			resourceForPostgres(*resource))
		transaction.Commit()
	} else {
		newVersion, _ := semver.NewVersion(resource.Version)
		existentVersion, _ := semver.NewVersion(existent)

		if newVersion.GreaterThan(existentVersion) {
			_, err = transaction.Exec(
				"UPDATE latest_security_resources SET raw = $1 WHERE raw @> jsonb_build_object('id', $2::text)",
				resourceForPostgres(*resource), resource.ID)
			transaction.Commit()
		}
	}

	availableVersions := r.retrieveAvailableVersions(resource.ID)
	_, err = r.db.Exec("UPDATE security_resources SET available_versions = $1 WHERE raw @> jsonb_build_object('id', $2::text) ", pq.Array(availableVersions), resource.ID)
	_, err = r.db.Exec("UPDATE latest_security_resources SET available_versions = $1 WHERE raw @> jsonb_build_object('id', $2::text) ", pq.Array(availableVersions), resource.ID)

	return err
}

func retrieveExistingVersion(transaction *sql.Tx, resource *Resource) string {
	var existent = ""
	transaction.QueryRow(`SELECT raw ->> 'version' AS version FROM latest_security_resources WHERE raw @> jsonb_build_object('id', $1::text)`, resource.ID).Scan(&existent)

	return existent
}

func (r *postgresRepository) retrieveAvailableVersions(id string) []string {
	var availableVersions = []string{}
	r.db.QueryRow(`SELECT ARRAY(SELECT raw ->> 'version' from security_resources WHERE raw ->> 'id' = $1::text) FROM security_resources WHERE raw ->> 'id' = $1::text LIMIT 1;`, id).Scan(pq.Array(&availableVersions))

	return availableVersions
}

func (r *postgresRepository) FindById(id string) (*Resource, error) {
	result := new(resourceForPostgres)
	availableVersions := []string{}
	err := r.db.QueryRow(`SELECT available_versions, raw FROM latest_security_resources WHERE raw @> jsonb_build_object('id', $1::text)`, id).Scan(pq.Array(&availableVersions), &result)

	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}

	result.AvailableVersions = availableVersions

	return (*Resource)(result), err
}

func (r *postgresRepository) FindAll() ([]*Resource, error) {
	rows, err := r.db.Query(`SELECT available_versions, raw FROM latest_security_resources`)
	defer rows.Close()

	var result []*Resource
	for rows.Next() {
		availableVersions := []string{}
		current := new(resourceForPostgres)
		if err = rows.Scan(pq.Array(&availableVersions), &current); err != nil {
			return nil, err
		}
		current.AvailableVersions = availableVersions
		result = append(result, (*Resource)(current))
	}

	return result, err
}

func (r *postgresRepository) FindByVersion(id string, version string) (*Resource, error) {
	result := new(resourceForPostgres)
	availableVersions := []string{}
	err := r.db.QueryRow(`SELECT available_versions, raw FROM security_resources WHERE raw @> jsonb_build_object('id', $1::text, 'version', $2::text)`, id, version).Scan(pq.Array(&availableVersions), &result)
	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}

	result.AvailableVersions = availableVersions
	return (*Resource)(result), err
}
