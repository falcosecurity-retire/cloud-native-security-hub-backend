package infrastructure

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

func GetResourcesFromPath(path string) ([]*resource.Resource, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	var resources []*resource.Resource

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			resource, err := getResourceFromFile(path)
			if err != nil {
				return err
			}
			resources = append(resources, &resource)
		}
		return nil
	})

	return resources, nil
}

func getResourceFromFile(path string) (resource resource.Resource, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return
	}

	err = yaml.NewDecoder(file).Decode(&resource)
	if err != nil {
		return
	}

	return
}
