package resource

import (
	"database/sql"
	"github.com/lib/pq"

	"github.com/Masterminds/semver"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Save(resource *Resource) error {
	transaction, err := r.db.Begin()
	resourceDTO := NewResourceDTO(resource)

	_, err = transaction.Exec("INSERT INTO security_resources(raw) VALUES($1)", resourceDTO)

	if existent := retrieveExistingVersion(transaction, resource); existent == "" {
		_, err = transaction.Exec(
			"INSERT INTO latest_security_resources(raw) VALUES($1)",
			resourceDTO)
		transaction.Commit()
	} else {
		newVersion, _ := semver.NewVersion(resource.Version)
		existentVersion, _ := semver.NewVersion(existent)

		if newVersion.GreaterThan(existentVersion) {
			_, err = transaction.Exec(
				"UPDATE latest_security_resources SET raw = $1 WHERE raw @> jsonb_build_object('id', $2::text, 'kind', $3::text)",
				resourceDTO,
				resource.ID.Slug(),
				resource.ID.Kind())
			transaction.Commit()
		}
	}

	availableVersions := r.retrieveAvailableVersions(resource.ID)

	_, err = r.db.Exec("UPDATE security_resources SET available_versions = $1 WHERE raw @> jsonb_build_object('id', $2::text, 'kind', $3::text) ", pq.Array(availableVersions), resource.ID.Slug(), resource.ID.Kind())
	_, err = r.db.Exec("UPDATE latest_security_resources SET available_versions = $1 WHERE raw @> jsonb_build_object('id', $2::text, 'kind', $3::text) ", pq.Array(availableVersions), resource.ID.Slug(), resource.ID.Kind())
	return err
}

func retrieveExistingVersion(transaction *sql.Tx, resource *Resource) string {
	var existent = ""
	transaction.QueryRow(`SELECT raw ->> 'version' AS version FROM latest_security_resources WHERE raw @> jsonb_build_object('id', $1::text, 'kind', $2::text)`,
		resource.ID.Slug(),
		resource.ID.Kind()).Scan(&existent)

	return existent
}

func (r *postgresRepository) retrieveAvailableVersions(id ResourceID) []string {
	var availableVersions = []string{}
	r.db.QueryRow(`SELECT ARRAY(SELECT raw ->> 'version' from security_resources WHERE raw @> jsonb_build_object('id', $1::text, 'kind', $2::text) ORDER BY raw ->> 'version' DESC) FROM security_resources WHERE raw @> jsonb_build_object('id', $1::text, 'kind', $2::text) LIMIT 1;`, id.Slug(), id.Kind()).Scan(pq.Array(&availableVersions))

	return availableVersions
}

func (r *postgresRepository) FindById(id ResourceID) (*Resource, error) {
	result := new(ResourceDTO)
	availableVersions := []string{}
	err := r.db.QueryRow(
		`SELECT available_versions, raw FROM latest_security_resources WHERE raw @> jsonb_build_object('id', $1::text, 'kind', $2::text)`,
		id.Slug(),
		id.Kind()).Scan(pq.Array(&availableVersions), &result)

	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}

	result.AvailableVersions = availableVersions

	return result.ToEntity(), err
}

func (r *postgresRepository) FindAll() ([]*Resource, error) {
	rows, err := r.db.Query(`SELECT available_versions, raw FROM latest_security_resources`)
	defer rows.Close()

	var result []*Resource
	for rows.Next() {
		availableVersions := []string{}
		current := new(ResourceDTO)
		if err = rows.Scan(pq.Array(&availableVersions), &current); err != nil {
			return nil, err
		}
		current.AvailableVersions = availableVersions
		result = append(result, current.ToEntity())
	}

	return result, err
}

func (r *postgresRepository) FindByVersion(id ResourceID, version string) (*Resource, error) {
	result := new(ResourceDTO)
	availableVersions := []string{}
	err := r.db.QueryRow(`SELECT available_versions, raw FROM security_resources WHERE raw @> jsonb_build_object('id', $1::text, 'kind', $2::text, 'version', $3::text)`, id.Slug(), id.Kind(), version).Scan(pq.Array(&availableVersions), &result)
	if err == sql.ErrNoRows {
		return nil, ErrResourceNotFound
	}

	result.AvailableVersions = availableVersions
	return result.ToEntity(), err
}
