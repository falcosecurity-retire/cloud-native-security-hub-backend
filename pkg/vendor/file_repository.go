package vendor

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
	path                   string
	vendors                []*Vendor
	vendorsCacheFilledOnce sync.Once
	vendorsCacheError      error
}

func FromPath(path string) (*fileRepository, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	return &fileRepository{path: path}, nil
}

func (f *fileRepository) FindAll() (vendors []*Vendor, err error) {
	f.vendorsCacheFilledOnce.Do(f.fillVendorsCache)
	return f.vendors, f.vendorsCacheError
}

func (f *fileRepository) FindById(id string) (*Vendor, error) {
	f.vendorsCacheFilledOnce.Do(f.fillVendorsCache)
	idToFind := strings.ToLower(id)

	if f.vendorsCacheError != nil {
		return nil, f.vendorsCacheError
	}

	if len(f.vendors) == 0 {
		return nil, fmt.Errorf("no vendors")
	}

	for _, vendor := range f.vendors {
		if vendor.ID == idToFind {
			return vendor, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func vendorFromFile(path string) (vendor Vendor, err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	defer file.Close()
	if err != nil {
		return
	}

	err = yaml.NewDecoder(file).Decode(&vendor)
	if err != nil {
		return
	}

	return
}

func (f *fileRepository) fillVendorsCache() {
	var vendors []*Vendor
	f.vendorsCacheError = filepath.Walk(f.path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			vendor, err := vendorFromFile(path)
			if err != nil {
				return err
			}
			vendors = append(vendors, &vendor)
		}
		return nil
	})
	f.vendors = vendors
}

func (f *fileRepository) Save(vendor *Vendor) error {
	return errors.New("Operation not supported")
}
