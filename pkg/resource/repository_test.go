package resource

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileRepositoryWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../test/fixtures"
	repository := FileRepository{
		Path: path,
	}

	resources, _ := repository.All()

	resource := Resource{
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
		},
	}

	assert.Equal(t, []Resource{resource}, resources)
}
