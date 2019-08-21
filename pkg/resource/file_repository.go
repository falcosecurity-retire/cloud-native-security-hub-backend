package resource

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type file struct {
	Path string
}

func NewFile(path string) (*file, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return &file{path }, nil
}

func (f *file) All() (resources []Resource, err error) {
	err = filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			resource, err := resourceFromFile(path)
			if err != nil {
				return err
			}
			resources = append(resources, resource)
		}
		return nil
	})
	return
}

func resourceFromFile(path string) (resource Resource, err error) {
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
