package usecases

import (
	"github.com/falcosecurity/cloud-native-security-hub/pkg/resource"
)

type IncrementDownloadCount struct {
	Updater            resource.Updater
	ResourceID         string
}

func (c *IncrementDownloadCount) Execute() error {
	return c.Updater.IncrementDownloadCountFor(c.ResourceID)
}
