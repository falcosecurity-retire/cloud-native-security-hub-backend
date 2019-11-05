package resource

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type fileRepository struct {
	path                     string
	resourcesCache           []*Resource
	resourcesCacheFilledOnce sync.Once
	resourcesCacheError      error
}

func FromPath(path string) (*fileRepository, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return &fileRepository{path: path}, nil
}

func (f *fileRepository) FindAll() (resources []*Resource, err error) {
	f.resourcesCacheFilledOnce.Do(f.fillResourcesCache)
	return f.resourcesCache, f.resourcesCacheError
}

func (f *fileRepository) FindById(id string) (res *Resource, err error) {
	f.resourcesCacheFilledOnce.Do(f.fillResourcesCache)
	idToFind := strings.ToLower(id)

	if f.resourcesCacheError != nil {
		return nil, f.resourcesCacheError
	}

	if len(f.resourcesCache) == 0 {
		return nil, fmt.Errorf("no resources")
	}

	for _, resource := range f.resourcesCache {
		if resource.ID == idToFind {
			res = resource
			return
		}
	}

	err = fmt.Errorf("not found")
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

func (f *fileRepository) fillResourcesCache() {
	var resources []*Resource
	f.resourcesCacheError = filepath.Walk(f.path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			resource, err := resourceFromFile(path)
			if err != nil {
				return err
			}
			resources = append(resources, &resource)
		}
		return nil
	})
	f.resourcesCache = resources
}

func (f *fileRepository) Save(resources *Resource) error {
	return errors.New("Operation not supported")
}
