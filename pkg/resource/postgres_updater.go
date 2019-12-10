package resource

import "database/sql"

type postgresUpdater struct {
	db *sql.DB
}

func (u *postgresUpdater) IncrementDownloadCountFor(resourceID string) error {
	transaction, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = transaction.Exec("UPDATE security_resources SET download_count = download_count + 1 WHERE raw ->> 'id' = $1", resourceID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	_, err = transaction.Exec("UPDATE latest_security_resources SET download_count = download_count + 1 WHERE raw ->> 'id' = $1", resourceID)
	if err != nil {
		transaction.Rollback()
		return err
	}

	err = transaction.Commit()
	return err
}

func NewPostgresUpdater(db *sql.DB) Updater {
	return &postgresUpdater{db: db}
}
