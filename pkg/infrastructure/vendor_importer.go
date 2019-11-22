package infrastructure

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"

	"github.com/falcosecurity/cloud-native-security-hub/pkg/vendor"
)

func GetVendorsFromPath(path string) ([]*vendor.Vendor, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	var vendors []*vendor.Vendor

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".yaml" {
			vendor, err := getVendorFromFile(path)
			if err != nil {
				return err
			}
			vendors = append(vendors, &vendor)
		}
		return nil
	})

	return vendors, nil
}

func getVendorFromFile(path string) (vendor vendor.Vendor, err error) {
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
