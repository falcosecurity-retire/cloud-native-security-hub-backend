package resource

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"os"
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
	assert.Equal(t, nil, err)
	assert.Equal(t, []Resource{
		{Name: "foo"},
		{Name: "bar"},
	}, resources)
}

type fileRepository struct {
	filenames []string
}

func (f *fileRepository) All() (resources []Resource, err error) {
	for _, filename := range f.filenames {
		file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
		if err != nil {
			break
		}
		var resource Resource
		err = yaml.NewDecoder(file).Decode(&resource)
		if err != nil {
			break
		}
		resources = append(resources, resource)
		file.Close()
	}
	return
}

func TestFileRepository(t *testing.T) {
	filenames := []string{"../../test/falco_resource.yaml"}
	repository := fileRepository{filenames}
	resources, err := repository.All()
	assert.Equal(t, nil, err)
	assert.Equal(t, len(filenames), len(resources))

	falcoRule := &FalcoRule{
		ApiVersion:  "v1",
		Kind:        "FalcoRules",
		Vendor:      "foo",
		Name:        "FooBar",
		Description: "FooBar Description\n",
		Readme:      "README.md",
		Keywords:    []string{"monitoring", "security", "cryptomining"},
		Icon:        "https://sysdig.com/image.png",
		Maintainers: []*Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
			{
				Name:  "nestorsalceda",
				Email: "nestor.salceda@sysdig.com",
			},
		},
		Rules: []*FalcoRuleData{
			{
				Raw: `- list: my_programs
  items: [ls, cat, pwd]

- rule: my_programs_opened_file
  desc: track whenever a set of programs opens a file
  condition: proc.name in (my_programs) and evt.type=open
  output: a tracked program opened a file (user=%user.name command=%proc.cmdline file=%fd.name)
  priority: INFO
`,
			},
			{
				File: "rules.yaml",
			},
			{
				File: "rules2.yaml",
			},
		},
	}

	assert.Equal(t, falcoRule, resources[0].ToFalcoRule())
}
