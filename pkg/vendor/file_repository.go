package vendor

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	Path string
}

func NewFile(path string) (*file, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return &file{path}, nil
}

func (f *file) FindAll() (resources []*Resource, err error) {
	err = filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			resource, err := resourceFromFile(path)
			if err != nil {
				return err
			}
			resources = append(resources, &resource)
		}
		return nil
	})
	return
}

func (f *file) FindById(id string) (res *Resource, err error) {
	idToFind := strings.ToLower(id)
	err = filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			resource, err := resourceFromFile(path)
			if err != nil {
				return err
			}
			resourceHash := strings.ToLower(resource.Hash())
			resourceName := strings.ToLower(resource.Name)
			if resourceHash == idToFind || resourceName == idToFind {
				res = &resource
				return nil
			}
		}
		return nil
	})
	if err == nil && res == nil {
		err = fmt.Errorf("not found")
	}
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
