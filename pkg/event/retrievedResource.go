package event

import "github.com/falcosecurity/cloud-native-security-hub/pkg/resource"

type RetrievedResource struct {
	ResourceID string
	Updater    resource.Updater
}

func (r *RetrievedResource) Handle() error {
	return r.Updater.IncrementDownloadCountFor(r.ResourceID)
}
