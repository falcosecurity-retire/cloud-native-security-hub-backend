package resource

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type file struct {
	path                     string
	resourcesCache           []*Resource
	resourcesCacheFilledOnce sync.Once
	resourcesCacheError      error
}

func FromPath(path string) (*file, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return &file{path: path}, nil
}

func (f *file) FindAll() (resources []*Resource, err error) {
	f.resourcesCacheFilledOnce.Do(f.fillResourcesCache)
	return f.resourcesCache, f.resourcesCacheError
}

func (f *file) FindById(id string) (res *Resource, err error) {
	f.resourcesCacheFilledOnce.Do(f.fillResourcesCache)
	idToFind := strings.ToLower(id)

	if f.resourcesCacheError != nil {
		return nil, f.resourcesCacheError
	}

	if len(f.resourcesCache) == 0 {
		return nil, fmt.Errorf("no resources")
	}

	for _, resource := range f.resourcesCache {
		resourceHash := strings.ToLower(resource.Hash())
		if resourceHash == idToFind {
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

func (f *file) fillResourcesCache() {
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
