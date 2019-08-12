package repository

import (
	"cloud-native-visibility-hub/pkg/resource"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestFileRepositoryWalksADirectoryAndExtractResources(t *testing.T) {
	path := "../../../test/fixtures"
	fileRepository := File{Path: path}
	resources, _ := fileRepository.All()

	files, err := ioutil.ReadDir(path)
	assert.Equal(t, nil, err)

	numYamls := 0
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".yaml") {
			numYamls++
		}
	}

	assert.Equal(t, numYamls, len(resources))

	res := resource.Resource{
		ApiVersion:  "v1",
		Kind:        "FalcoRules",
		Vendor:      "foo",
		Name:        "FooBar",
		Description: "FooBar Description\n",
		Keywords:    []string{"monitoring", "security", "cryptomining"},
		Icon:        "https://sysdig.com/image.png",
		Maintainers: []*resource.Maintainer{
			{
				Name:  "bencer",
				Email: "bencer@sysdig.com",
			},
			{
				Name:  "nestorsalceda",
				Email: "nestor.salceda@sysdig.com",
			},
		},
		Rules: []*resource.FalcoRuleData{
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

//	assert.Equal(t, []resource.Resource{res}, resources)
	_ = res
}
