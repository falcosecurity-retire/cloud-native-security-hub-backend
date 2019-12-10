//go:generate mockgen -destination mock_resource/mock_updater.go github.com/falcosecurity/cloud-native-security-hub/pkg/resource Updater
package resource

type Updater interface {
	IncrementDownloadCountFor(resourceID string) error
}
