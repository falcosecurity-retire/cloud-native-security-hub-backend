package resource

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Repository interface {
	All() ([]Resource, error)
}

type FileRepository struct {
	Path string
}

func (f *FileRepository) All() ([]Resource, error) {
	resources := []Resource{}
	err := filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
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
	return resources, err
}

func (f *FileRepository) resourceFromFile(path string) (Resource, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return Resource{}, err
	}

	var resource Resource
	err = yaml.NewDecoder(file).Decode(&resource)
	if err != nil {
		return Resource{}, err
	}

	return resource, nil
}
