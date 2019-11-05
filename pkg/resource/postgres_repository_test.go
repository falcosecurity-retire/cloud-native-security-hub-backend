package resource

import (
	"database/sql"
	"os"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAResource(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := NewPostgresRepository(db)

	resource := apacheResource()
	repository.Save(resource)

	retrieved, _ := repository.FindById("apache")
	assert.Equal(t, resource, retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
}

func apacheResource() *Resource {
	return &Resource{
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
}

func mongodbResource() *Resource {
	return &Resource{
		Kind:        "FalcoRules",
		Vendor:      "Mongo",
		ID:          "mongodb",
		Name:        "MongoDB",
		Description: "# MongoDB Falco Rules\n",
		Keywords:    []string{"database"},
		Icon:        "https://upload.wikimedia.org/wikipedia/en/thumb/4/45/MongoDB-Logo.svg/2560px-MongoDB-Logo.svg.png",
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
				Raw: "- macro: mongo_consider_syscalls\n  condition: (evt.num < 0)",
			},
		},
	}
}

func TestFindAllResources(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	repository := NewPostgresRepository(db)

	repository.Save(apacheResource())
	repository.Save(mongodbResource())

	retrieved, _ := repository.FindAll()
	assert.Equal(t, []*Resource{apacheResource(), mongodbResource()}, retrieved)

	db.Exec("TRUNCATE TABLE security_resources")
}
