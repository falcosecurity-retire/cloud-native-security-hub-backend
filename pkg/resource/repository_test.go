package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyRepository struct {
}

func (d *dummyRepository) All() ([]Resource, error) {
	resources := []Resource{
		{Name: "foo"},
		{Name: "bar"},
	}
	return resources, nil
}

func TestRepository(t *testing.T) {
	t.Parallel()

	dummyRepo := dummyRepository{}
	resources, err := dummyRepo.All()
	assert.Equal(t, err, nil)
	assert.Equal(t, resources, []Resource{
		{Name: "foo"},
		{Name: "bar"},
	})
}