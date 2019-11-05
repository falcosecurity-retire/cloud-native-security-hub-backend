package resource

import (
	"database/sql"
	"os"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAResource(t *testing.T) {
	resource := &Resource{
		ID:          "apache",
		Kind:        "FalcoRules",
		Version:     "1.0.0",
		Vendor:      "Apache",
		Name:        "Apache",
		Description: "# Apache Falco Rules\n",
		Keywords:    []string{"web"},
		Icon:        "https://upload.wikimedia.org/wikipedia/commons/thumb/d/db/Apache_HTTP_server_logo_%282016%29.svg/300px-Apache_HTTP_server_logo_%282016%29.svg.png",
		Maintainers: []*Maintainer{
			{
				Name: "nestorsalceda",
				Link: "github.com/nestorsalceda",
			},
			{
				Name: "fedebarcelona",
				Link: "github.com/tembleking",
			},
		},
		Rules: []*FalcoRuleData{
			{
				Raw: "- macro: apache_consider_syscalls\n  condition: (evt.num < 0)",
			},
		},
	}

	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := PostgresRepository{db: db}

	repository.Save(resource)

	retrieved, _ := repository.FindById("apache")
	assert.Equal(t, resource, retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
}
