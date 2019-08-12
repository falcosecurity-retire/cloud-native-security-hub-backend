package repository

import (
	"cloud-native-visibility-hub/pkg/resource"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type File struct {
	Path string
}

func (f *File) All() (resources []resource.Resource, err error) {
	err = filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".yaml" {
			resource, err := f.resourceFromFile(path)
			if err != nil {
				return err
			}
			resources = append(resources, resource)
		}
		return nil
	})
	return
}

func (f *File) resourceFromFile(path string) (resource resource.Resource, err error) {
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
